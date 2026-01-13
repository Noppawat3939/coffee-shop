package controllers

import (
	"backend/dto"
	"backend/pkg/types"
	"backend/repository"
	"backend/services"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentController struct {
	repo    repository.OrderRepo
	service services.PaymentService
}

func NewPaymentController(r repository.OrderRepo, s services.PaymentService) *paymentController {
	return &paymentController{r, s}
}

func (pc *paymentController) CreatePaymentTransactionLog(c *gin.Context) {
	var req dto.CreateTxnLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	res, err := pc.service.CreatePaymentTransactionLog(req)

	if err != nil {
		util.Error(c, http.StatusConflict, "failed create payment transaction log")
		return
	}

	util.Success(c, res)
}

func (pc *paymentController) EnquiryPayment(c *gin.Context) {
	var req dto.EnquireTxnRequst

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	filter := types.Filter{
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

	res := dto.EnquireTxnResponse{
		TransactionNumber: log.TransactionNumber,
		Amount:            log.Amount,
		Status:            log.Status,
		ExpiredAt:         log.ExpiredAt,
		CreatedAt:         log.CreatedAt,
		Order: dto.EnquireTxnWithOrderResponse{
			ID:          log.Order.ID,
			OrderNumber: log.Order.OrderNumber,
			Status:      log.Order.Status,
			Total:       log.Order.Total,
		},
	}

	util.Success(c, res)
}
