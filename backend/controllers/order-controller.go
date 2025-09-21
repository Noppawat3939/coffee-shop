package controllers

import (
	"backend/dto"
	hlp "backend/helpers"
	"backend/models"
	"backend/repository"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderController struct {
	repo repository.OrderRepo
	db   *gorm.DB
}

var OrderStatus = struct {
	ToPay     string
	Paid      string
	Cancelled string
}{ToPay: "to_pay", Paid: "paid", Cancelled: "cancelled"}

func NewOrderController(repo repository.OrderRepo, db *gorm.DB) *orderController {
	return &orderController{repo, db}
}

func (oc *orderController) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(c)
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
				invalidMenuVariationIDs = append(invalidMenuVariationIDs, hlp.IntToString(int(v.MenuVariationID)))
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

		// return error if any invalid menu_variations
		if len(invalidMenuVariationIDs) > 0 {
			idStr := strings.Join(invalidMenuVariationIDs, ",")
			errStatus = http.StatusNotAcceptable
			errMsg = "invalid menu_variation_ids: " + idStr

			return fmt.Errorf("invalid menu_variation_ids: %s", idStr)
		}

		// create orders
		order = models.Order{
			OrderNumber: uuid.NewString(),
			Status:      OrderStatus.ToPay,
			Customer:    customer,
			Total:       total,
		}

		if _, err := oc.repo.CreateOrder(order, tx); err != nil {
			return err
		}

		// create order_menu_variations
		for i := range odVariations {
			odVariations[i].OrderID = order.ID
			if _, err := oc.repo.CreateOrderMenuVariation(odVariations[i], tx); err != nil {
				return err
			}
		}

		// create order_status_logs
		if _, err := oc.repo.CreateOrderStatusLog(models.OrderStatusLog{
			OrderID: order.ID,
			Status:  order.Status,
		}, tx); err != nil {
			return err
		}

		return tx.First(&order, order.ID).Error
	})

	if err != nil {
		hlp.Error(c, errStatus, errMsg)
		return
	}

	hlp.Success(c, order)
}

func (oc *orderController) GetOrderByID(c *gin.Context) {
	id := hlp.ParamToInt(c, "id")

	order, err := oc.repo.FindOneOrder(id)
	if err != nil {
		hlp.ErrorNotFound(c)

		return
	}

	hlp.Success(c, order)
}
