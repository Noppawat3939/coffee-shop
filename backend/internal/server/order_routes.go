package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialOrderRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	odSvc := service.NewOrderService(repo)
	controller := controllers.NewOrderController(repo, odSvc, cfg.DB)

	order := r.Group("/Orders")
	{
		order.POST("", middleware.AuthGuard(), controller.CreateOrder)
		order.GET("", controller.GetOrders)
		order.GET("/:id", middleware.AuthGuard(), controller.GetOrderByID)
		order.GET("/order-number/:order_number", middleware.AuthGuard(), controller.GetOrderByOrderNumber)
	}
}
