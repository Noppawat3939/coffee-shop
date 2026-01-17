package controllers

import (
	"backend/dto"
	"backend/repository"
	"backend/services"
	"backend/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentController struct {
	repo        repository.OrderRepo
	paymentRepo repository.PaymentRepo
	service     services.PaymentService
	db          *gorm.DB
}

func NewPaymentController(r repository.OrderRepo, pr repository.PaymentRepo, s services.PaymentService, d *gorm.DB) *paymentController {
	return &paymentController{r, pr, s, d}
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

	q := map[string]interface{}{
		"transaction_number": req.TransactionNumber,
	}

	if req.Status != "" {
		q["status"] = req.Status
	}

	res, err := pc.service.FindOnePaymentLog(q)

	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, res)
}

func (pc *paymentController) UpdatePaymentAndOrderStatus(c *gin.Context, status string) {
	ref := c.Param("order_number")
	err := pc.db.Transaction(func(tx *gorm.DB) error {
		_, err := pc.service.UpdatePaymentStatus(ref, status, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		util.Error(c, http.StatusNotFound, fmt.Sprintf("order number %s already status %s", ref, status))
		return
	}

	util.Success(c)
}
