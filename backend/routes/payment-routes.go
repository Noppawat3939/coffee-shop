package routes

import (
	ctl "backend/controllers"
	"backend/middleware"
	"backend/models"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	odRepo := repository.NewOrderRepository(cfg.DB)
	payRepo := repository.NewPaymentRepository(cfg.DB)
	svc := services.NewPaymentService(odRepo, payRepo)
	odSvc := services.NewOrderService(odRepo)
	controller := ctl.NewPaymentController(payRepo, svc, odSvc, cfg.DB)

	payment := r.Group("/Payments/txns", middleware.AuthGuard())
	{
		payment.GET("/", controller.GetPaymentTransactions)
		payment.POST("/order", middleware.IdempotencyMiddleware(cfg.DB), controller.CreatePaymentTransactionLog)
		payment.POST("/enquiry", controller.EnquiryPayment)
		payment.POST("/:order_number/paid", middleware.IdempotencyMiddleware(cfg.DB), func(ctx *gin.Context) { controller.UpdatePaymentAndOrderStatus(ctx, models.OrderStatus.Paid) })
		payment.POST("/:order_number/canceled", middleware.IdempotencyMiddleware(cfg.DB), func(ctx *gin.Context) { controller.UpdatePaymentAndOrderStatus(ctx, models.OrderStatus.Canceled) })
	}
}
