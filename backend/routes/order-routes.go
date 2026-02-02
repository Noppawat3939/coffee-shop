package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialOrderRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	odSvc := services.NewOrderService(repo)
	controller := controllers.NewOrderController(repo, odSvc, cfg.DB)

	order := r.Group("/Orders")
	{
		order.POST("", middleware.AuthGuard(), middleware.AuthGuard(), controller.CreateOrder)
		order.GET("", controller.GetOrders)
		order.GET("/:id", middleware.AuthGuard(), controller.GetOrderByID)
		order.GET("/order-number/:order_number", middleware.AuthGuard(), controller.GetOrderByOrderNumber)
	}
}
