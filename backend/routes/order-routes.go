package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialOrderRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewOrderRepository(db)
	controller := controllers.NewOrderController(repo, db)

	order := r.Group("/Orders")
	{
		order.POST("", controller.CreateOrder)
		order.GET("/:id", controller.GetOrderByID)
		order.GET("/order-number/:order_number", controller.GetOrderByOrderNumber)
		order.PATCH("/:id/paid", func(ctx *gin.Context) {
			controller.UpdateOrderStatus(ctx, controllers.OrderStatus.Paid)
		})
		order.PATCH("/:id/canceled", func(ctx *gin.Context) {
			controller.UpdateOrderStatus(ctx, controllers.OrderStatus.Canceled)
		})
	}
}
