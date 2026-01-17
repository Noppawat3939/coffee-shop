package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialOrderRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	controller := controllers.NewOrderController(repo, cfg.DB)

	order := r.Group("/Orders")
	{
		order.POST("", controller.CreateOrder)
		order.GET("", controller.GetOrders)
		order.GET("/:id", controller.GetOrderByID)
		order.GET("/order-number/:order_number", controller.GetOrderByOrderNumber)
		// order.PATCH("/:id/paid", func(ctx *gin.Context) {
		// 	controller.UpdateOrderStatus(ctx, models.OrderStatus.Paid)
		// })
		// order.PATCH("/:id/canceled", func(ctx *gin.Context) {
		// 	controller.UpdateOrderStatus(ctx, models.OrderStatus.Canceled)
		// })
	}
}
