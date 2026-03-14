package handler

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/pagination"
	"backend/pkg/response"
	"backend/pkg/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderHandler struct {
	repo  repository.OrderRepo
	odSvc service.OrderService
	db    *gorm.DB
}

func NewOrderHandler(repo repository.OrderRepo, odSvc service.OrderService, db *gorm.DB) *orderHandler {
	return &orderHandler{repo, odSvc, db}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	user := auth.GetUserFromContext(c)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	customer := "guest"
	if req.Customer != nil && *req.Customer != "" {
		customer = *req.Customer
	}

	var order model.Order
	var total float64
	var odVariations []model.OrderMenuVariation
	var invalidMenuVariationIDs []string
	var errStatus int = http.StatusConflict
	var errMsg string = "failed create order"

	err := h.db.Transaction(func(tx *gorm.DB) error {
		// calculate total and prepare variations
		for _, v := range req.Variations {
			// check order variation id invalid
			mv, err := h.repo.FindOneMenuVariation(int(v.MenuVariationID))
			if err != nil {
				invalidMenuVariationIDs = append(invalidMenuVariationIDs, util.IntToString(int(v.MenuVariationID)))
				continue
			}

			total += float64(v.Amount) * mv.Price

			// append to order variations
			odVariations = append(odVariations, model.OrderMenuVariation{
				MenuVariationID: v.MenuVariationID,
				Amount:          v.Amount,
				Price:           mv.Price,
			})
		}

		if len(invalidMenuVariationIDs) > 0 {
			ids := strings.Join(invalidMenuVariationIDs, ",")
			errStatus = http.StatusNotAcceptable
			errMsg = "invalid menu_variation_ids: " + ids

			return fmt.Errorf("invalid menu_variation_ids: %s", ids)
		}

		// create orders
		order = model.Order{
			OrderNumber: uuid.NewString(),
			Status:      model.OrderStatus.ToPay,
			Customer:    customer,
			Total:       total,
			EmployeeID:  &user.ID,
			MemberID:    &req.MemberID,
		}

		if _, err := h.repo.CreateOrder(&order, tx); err != nil {
			return err
		}

		if _, err := h.odSvc.CreateMenuVariations(odVariations, order, tx); err != nil {
			return err
		}

		if _, err := h.odSvc.CreateLog(order, tx); err != nil {
			return err
		}

		return tx.First(&order, order.ID).Error
	})

	if err != nil {
		response.Error(c, errStatus, errMsg)
		return
	}

	response.Success(c, order)
}

func (h *orderHandler) GetOrderByID(c *gin.Context) {
	id := util.ToInt(c.Param("id"))

	order, err := h.repo.FindOneOrder(id)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, order)
}

func (h *orderHandler) GetOrderByOrderNumber(c *gin.Context) {
	order_number := c.Param("order_number")

	order, err := h.repo.FindOneOrderByOrderNumber(order_number)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, order)
}

func (h *orderHandler) GetOrders(c *gin.Context) {
	status := c.Param("status")
	id := util.ToInt((c.Param("id")))
	p := pagination.NewFromQuery(c)

	q := map[string]interface{}{
		"id":     id,
		"status": status,
	}

	orders, err := h.repo.FindAllOrders(q, p.Page, p.Limit)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, orders)
}
