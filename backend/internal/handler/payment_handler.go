package handler

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/pagination"
	"backend/pkg/response"
	"backend/pkg/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentHandler struct {
	paymentRepo repository.PaymentRepo
	paymentSvc  service.PaymentService
	pointSvc    service.MemberPointService
	odSvc       service.OrderService
	db          *gorm.DB
}

func NewPaymentHandler(paymentRepo repository.PaymentRepo, paymentSvc service.PaymentService, pointSvc service.MemberPointService, odSvc service.OrderService, db *gorm.DB) *paymentHandler {
	return &paymentHandler{paymentRepo, paymentSvc, pointSvc, odSvc, db}
}

func (h *paymentHandler) CreatePaymentTransactionLog(c *gin.Context) {
	var req dto.CreateTxnLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	res, err := h.paymentSvc.CreatePaymentTransactionLog(req)

	if err != nil {
		response.Error(c, http.StatusConflict, "failed create payment transaction log")
		return
	}

	response.Success(c, res)
}

func (h *paymentHandler) EnquiryPayment(c *gin.Context) {
	var req dto.EnquireTxnRequst

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	q := map[string]interface{}{
		"transaction_number": req.TransactionNumber,
	}

	if req.Status != "" {
		q["status"] = req.Status
	}

	res, err := h.paymentSvc.FindOnePaymentLog(q)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, res)
}

func (h *paymentHandler) UpdatePaymentAndOrderStatus(c *gin.Context, status string) {
	odNo := c.Param("order_number")

	err := h.db.Transaction(func(tx *gorm.DB) error {
		// update payment log not expired
		_, err := h.paymentSvc.UpdatePaymentStatus(odNo, status, tx)
		if err != nil {
			return err
		}

		// update order and create order_status_log
		order, err := h.odSvc.UpdateOrderStatusAndLog(odNo, status, tx)
		if err != nil {
			return err
		}

		// update point earn
		if status == model.OrderStatus.Paid {
			if err := h.pointSvc.EarnPointFromOrder(order, tx); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		response.Error(c, http.StatusNotFound, fmt.Sprintf("order number %s already status %s", odNo, status))
		return
	}

	response.Success(c)
}

func (h *paymentHandler) GetPaymentTransactions(c *gin.Context) {
	p := pagination.NewFromQuery(c)
	idStr := c.Query("id")
	status := c.Query("status")
	transaction_number := c.Query("transaction_number")
	order_number_ref := c.Query("order_number_ref")

	q := make(map[string]interface{})

	if status != "" {
		q["status"] = status
	}
	if idStr != "" {
		q["id"] = util.ToInt(idStr)
	}
	if transaction_number != "" {
		q["transaction_number"] = transaction_number
	}
	if order_number_ref != "" {
		q["order_number_ref"] = order_number_ref
	}

	logs, err := h.paymentRepo.FindAllTransactions(q, p.Page, p.Limit)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, logs)
}
