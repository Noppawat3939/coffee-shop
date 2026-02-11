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
	paymentRepo repository.PaymentRepo
	paymentSvc  services.PaymentService
	odSvc       services.OrderService
	db          *gorm.DB
}

func NewPaymentController(paymentRepo repository.PaymentRepo, paymentSvc services.PaymentService, odSvc services.OrderService, db *gorm.DB) *paymentController {
	return &paymentController{paymentRepo, paymentSvc, odSvc, db}
}

func (pc *paymentController) CreatePaymentTransactionLog(c *gin.Context) {
	var req dto.CreateTxnLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
	}

	res, err := pc.paymentSvc.CreatePaymentTransactionLog(req)

	if err != nil {
		util.Error(c, http.StatusConflict, "failed create payment transaction log")
	}

	util.Success(c, res)
}

func (pc *paymentController) EnquiryPayment(c *gin.Context) {
	var req dto.EnquireTxnRequst

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
	}

	q := map[string]interface{}{
		"transaction_number": req.TransactionNumber,
	}

	if req.Status != "" {
		q["status"] = req.Status
	}

	res, err := pc.paymentSvc.FindOnePaymentLog(q)

	if err != nil {
		util.ErrorNotFound(c)
	}

	util.Success(c, res)
}

func (pc *paymentController) UpdatePaymentAndOrderStatus(c *gin.Context, status string) {
	odNo := c.Param("order_number")
	err := pc.db.Transaction(func(tx *gorm.DB) error {
		// update payment log not expired
		_, err := pc.paymentSvc.UpdatePaymentStatus(odNo, status, tx)
		if err != nil {
			return err
		}

		// update order and create order_status_log
		if _, err := pc.odSvc.UpdateOrderStatusAndLog(odNo, status, tx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		util.Error(c, http.StatusNotFound, fmt.Sprintf("order number %s already status %s", odNo, status))
	}

	util.Success(c)
}

func (pc *paymentController) GetPaymentTransactions(c *gin.Context) {
	page, limit := util.BuildPagination(c)
	idStr := c.Param("id")
	status := c.Param("status")
	transaction_number := c.Param("transaction_number")
	order_number_ref := c.Param("order_number_ref")

	var id *int
	if idStr != "" {
		v := util.ParamToInt(c, "id")
		id = &v
	}

	q := util.CleanNilMap(map[string]interface{}{
		"id":                 id,
		"status":             status,
		"transaction_number": transaction_number,
		"order_number_ref":   order_number_ref,
	})

	logs, err := pc.paymentRepo.FindAllTransactions(q, page, limit)

	if err != nil {
		util.ErrorNotFound(c)
	}

	util.Success(c, logs)
}
