package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/pkg/types"
	"backend/repository"
	"backend/util"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderController struct {
	repo repository.OrderRepo
	db   *gorm.DB
}

func NewOrderController(repo repository.OrderRepo, db *gorm.DB) *orderController {
	return &orderController{repo, db}
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
			Status:      models.OrderStatus.ToPay,
			Customer:    customer,
			Total:       total,
			EmployeeID:  user.ID,
		}

		if _, err := oc.repo.CreateOrder(&order, tx); err != nil {
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

func (oc *orderController) UpdateOrderStatus(c *gin.Context, statusToUpdate string) {
	id := util.ParamToInt(c, "id")

	order, err := oc.repo.FindOneOrder(id)
	if err != nil {
		util.ErrorNotFound(c)

		return
	}

	// check status not allowed to update
	allowed, ok := allowedUpdateStatus[order.Status]
	if !ok || !slices.Contains(allowed, statusToUpdate) {
		msg := fmt.Sprintf("%s %s", "current status not allowed to update to", statusToUpdate)
		util.Error(c, http.StatusNotAcceptable, msg)

		return
	}

	err = oc.db.Transaction(func(tx *gorm.DB) error {
		_, err := oc.repo.UpdateOrderByID(id, models.Order{
			Status: statusToUpdate,
		}, tx)

		if err != nil {
			return err
		}

		// update payment_transaction_log where by order_id and status is to_pay
		// update payment_transaction_log status to body.status
		filter := types.Filter{
			"order_id": id,
			"status":   models.OrderStatus.ToPay,
		}

		updateLog := models.PaymentOrderTransactionLog{
			Status: statusToUpdate,
		}

		if statusToUpdate == models.OrderStatus.Paid || statusToUpdate == models.OrderStatus.Canceled {
			if _, err := oc.repo.UpdatePaymentLog(filter, updateLog); err != nil {
				return err
			}
		}

		if _, err := oc.repo.CreateOrderStatusLog(models.OrderStatusLog{
			OrderID: order.ID,
			Status:  statusToUpdate,
		}, tx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		util.Error(c, http.StatusConflict, "failed update order status to paid")
		return
	}

	util.Success(c)
}

func (oc *orderController) GetOrders(c *gin.Context) {
	status := c.Param("status")
	id := util.ParamToInt(c, "id")
	page := util.ToInt(c.DefaultQuery("page", fmt.Sprint(util.DefaultPage)))
	limit := util.ToInt(c.DefaultQuery("limit", fmt.Sprint(util.DefaultLimit)))

	filter := types.Filter{
		"id":     id,
		"status": status,
	}

	orders, err := oc.repo.FindAllOrders(filter, page, limit)
	if err != nil {
		util.ErrorNotFound(c)

		return
	}

	util.Success(c, orders)
}

var allowedUpdateStatus = map[string][]string{
	models.OrderStatus.ToPay: {models.OrderStatus.Paid, models.OrderStatus.Canceled},
}
