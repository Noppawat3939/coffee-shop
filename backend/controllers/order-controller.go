package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/services"
	"backend/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderController struct {
	repo  repository.OrderRepo
	odSvc services.OrderService
	db    *gorm.DB
}

func NewOrderController(repo repository.OrderRepo, odSvc services.OrderService, db *gorm.DB) *orderController {
	return &orderController{repo, odSvc, db}
}

func (oc *orderController) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	user, ok := util.GetUserFromHeader(c)

	if !ok {
		util.ErrorUnauthorized(c)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	customer := "guest"
	if req.Customer != nil && *req.Customer != "" {
		customer = *req.Customer
	}

	var order models.Order
	var total float64
	var odVariations []models.OrderMenuVariation
	var invalidMenuVariationIDs []string
	var errStatus int = http.StatusConflict
	var errMsg string = "failed create order"

	err := oc.db.Transaction(func(tx *gorm.DB) error {
		// calculate total and prepare variations
		for _, v := range req.Variations {
			// check order variation id invalid
			mv, err := oc.repo.FindOneMenuVariation(int(v.MenuVariationID))
			if err != nil {
				invalidMenuVariationIDs = append(invalidMenuVariationIDs, util.IntToString(int(v.MenuVariationID)))
				continue
			}

			total += float64(v.Amount) * mv.Price

			// append to order variations
			odVariations = append(odVariations, models.OrderMenuVariation{
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
		order = models.Order{
			OrderNumber: uuid.NewString(),
			Status:      models.OrderStatus.ToPay,
			Customer:    customer,
			Total:       total,
			EmployeeID:  user.ID,
		}

		if _, err := oc.repo.CreateOrder(&order, tx); err != nil {
			return err
		}

		if _, err := oc.odSvc.CreateMenuVariations(odVariations, order, tx); err != nil {
			return err
		}

		if _, err := oc.odSvc.CreateLog(order, tx); err != nil {
			return err
		}

		return tx.First(&order, order.ID).Error
	})

	if err != nil {
		util.Error(c, errStatus, errMsg)
		return
	}

	util.Success(c, order)
}

func (oc *orderController) GetOrderByID(c *gin.Context) {
	id := util.ParamToInt(c, "id")

	order, err := oc.repo.FindOneOrder(id)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, order)
}

func (oc *orderController) GetOrderByOrderNumber(c *gin.Context) {
	order_number := c.Param("order_number")

	order, err := oc.repo.FindOneOrderByOrderNumber(order_number)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, order)
}

func (oc *orderController) GetOrders(c *gin.Context) {
	status := c.Param("status")
	id := util.ParamToInt(c, "id")
	page, limit := util.BuildPagination(c)

	q := map[string]interface{}{
		"id":     id,
		"status": status,
	}

	orders, err := oc.repo.FindAllOrders(q, page, limit)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, orders)
}
