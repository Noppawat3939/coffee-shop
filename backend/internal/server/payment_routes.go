package server

import (
	ctl "backend/internal/handler"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	odRepo := repository.NewOrderRepository(cfg.DB)
	payRepo := repository.NewPaymentRepository(cfg.DB)
	pointRepo := repository.NewMemberPointRepository(cfg.DB)
	svc := service.NewPaymentService(odRepo, payRepo)
	odSvc := service.NewOrderService(odRepo)
	pointSvc := service.NewMemberPointService(pointRepo)
	handler := ctl.NewPaymentHandler(payRepo, svc, pointSvc, odSvc, cfg.DB)

	payment := r.Group("/Payments/txns", middleware.AuthGuard())
	{
		payment.GET("", handler.GetPaymentTransactions)
		payment.POST("/order", middleware.IdempotencyMiddleware(cfg.DB), handler.CreatePaymentTransactionLog)
		payment.POST("/enquiry", handler.EnquiryPayment)
		payment.POST("/:order_number/paid", middleware.IdempotencyMiddleware(cfg.DB), func(ctx *gin.Context) { handler.UpdatePaymentAndOrderStatus(ctx, model.OrderStatus.Paid) })
		payment.POST("/:order_number/canceled", middleware.IdempotencyMiddleware(cfg.DB), func(ctx *gin.Context) { handler.UpdatePaymentAndOrderStatus(ctx, model.OrderStatus.Canceled) })
	}
}
