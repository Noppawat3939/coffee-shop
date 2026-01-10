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
		TransactionNumber: util.GenerateTransactionNumber(req.OrderNumber), // auto generate by uuid concat with order_id (unique)
		Status:            OrderStatus.ToPay,
		PaymentCode:       payload,
		QRSignature:       signature,
		ExpiredAt:         time.Now().Add(5 * time.Minute), // expired in 5 min
	}, nil)

	if logErr != nil {
		util.Error(c, http.StatusConflict, "failed create payment transaction log")
		return
	}

	res := dto.CreatePaymentTransactionLogResponse{
		TransactionNumber: log.TransactionNumber,
		Amount:            log.Amount,
		Status:            log.Status,
		PaymentCode:       log.PaymentCode,
		ExpiredAt:         log.ExpiredAt,
		CreatedAt:         log.CreatedAt,
	}

	util.Success(c, res)
}

func (pc *paymentController) EnquiryPayment(c *gin.Context) {
	var req dto.EnquirPaymentTransactionLogRequst

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	filter := map[string]interface{}{
		"transaction_number": req.TransactionNumber,
	}

	if req.Status != "" {
		filter["status"] = req.Status
	}

	log, err := pc.repo.FindOneTransaction(filter)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	res := dto.EnquiryPaymentTransactionLogResponse{
		TransactionNumber: log.TransactionNumber,
		Amount:            log.Amount,
		Status:            log.Status,
		ExpiredAt:         log.ExpiredAt,
		CreatedAt:         log.CreatedAt,
		Order: dto.EnquiryPaymentTransactionLogWithOrderResponse{
			ID:          log.Order.ID,
			OrderNumber: log.Order.OrderNumber,
			Status:      log.Order.Status,
			Total:       log.Order.Total,
		},
	}

	util.Success(c, res)
}
