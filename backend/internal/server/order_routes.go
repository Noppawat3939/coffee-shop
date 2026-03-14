package server

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialOrderRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	odSvc := service.NewOrderService(repo)
	handler := handler.NewOrderHandler(repo, odSvc, cfg.DB)

	order := r.Group("/Orders")
	{
		order.POST("", middleware.AuthGuard(), handler.CreateOrder)
		order.GET("", handler.GetOrders)
		order.GET("/:id", middleware.AuthGuard(), handler.GetOrderByID)
		order.GET("/order-number/:order_number", middleware.AuthGuard(), handler.GetOrderByOrderNumber)
	}
}
