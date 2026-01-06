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

	order, err := pc.repo.FindOneOrder(req.OrderID)

	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	log, logErr := pc.repo.CreatePaymentLog(models.PaymentOrderTransactionLog{
		OrderID:           uint(req.OrderID),
		Amount:            order.Total,
		TransactionNumber: generateTransactionNumber(req.OrderID), // auto generate by uuid concat with order_id (unique)
		Status:            OrderStatus.ToPay,
		PaymentCode:       "",                              // call server generate promptpay-qr
		ExpiredAt:         time.Now().Add(5 * time.Minute), // expired in 5 min
	}, nil)

	if logErr != nil {
		util.Error(c, http.StatusConflict, "failed create payment transaction log")
		return
	}

	util.Success(c, log)
}

func generateTransactionNumber(orderID int) string {
	return fmt.Sprintf("%s_%d", uuid.NewString(), orderID)
}
