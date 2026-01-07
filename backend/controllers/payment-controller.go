package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/services"
	"backend/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type paymentController struct {
	repo repository.OrderRepo
}

func NewPaymentController(repo repository.OrderRepo) *paymentController {
	return &paymentController{repo}
}

func (pc *paymentController) GeneratePromptPayQR(c *gin.Context) {
	var req dto.QRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	qr, err := services.GeneratePromptPayQR(req.Amount)
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "failed generate QR promptpay")
		return
	}

	util.Success(c, gin.H{"qr": qr})
}

func (pc *paymentController) CreatePaymentTransactionLog(c *gin.Context) {
	var req dto.CreatePaymentTransactionLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	order, err := pc.repo.FindOneOrderByOrderNumber(req.OrderNumber)

	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	// find prev log status is to_pay and auto expired
	prevLog, err := pc.repo.FindOneTransaction(
		map[string]interface{}{
			"order_id": order.ID,
			"status":   OrderStatus.ToPay,
		})

	if err == nil {
		_, updateErr := pc.repo.CancelAndExpirePaymentLogByID(int(prevLog.ID))

		if updateErr != nil {
			util.Error(c, http.StatusConflict, fmt.Sprintf("%s %s", "failed update payment_transaction_log status to", OrderStatus.Canceled))
			return
		}
	}

	payload, ppErr := services.GeneratePaymentCodePromptPayment(order.Total)
	if ppErr != nil {
		util.Error(c, http.StatusConflict, "failed generating payment_code")
		return
	}

	signature := services.SignPayload(payload)

	log, logErr := pc.repo.CreatePaymentLog(models.PaymentOrderTransactionLog{
		OrderID:           uint(order.ID),
		Amount:            order.Total,
		TransactionNumber: generateTransactionNumber(req.OrderNumber), // auto generate by uuid concat with order_id (unique)
		Status:            OrderStatus.ToPay,
		PaymentCode:       payload,
		QRSignature:       signature,
		ExpiredAt:         time.Now().Add(5 * time.Minute), // expired in 5 min
	}, nil)

	if logErr != nil {
		util.Error(c, http.StatusConflict, "failed create payment transaction log")
		return
	}

	result := map[string]interface{}{
		"transaction_number": log.TransactionNumber,
		"amount":             log.Amount,
		"status":             log.Status,
		"payment_code":       log.PaymentCode,
		"expired_at":         log.ExpiredAt,
		"created_at":         log.CreatedAt,
	}

	util.Success(c, result)
}

func generateTransactionNumber(orderNumber string) string {
	return fmt.Sprintf("%s_%s", uuid.NewString(), orderNumber)
}
