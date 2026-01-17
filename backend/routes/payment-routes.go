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
	controller := ctl.NewPaymentController(odRepo, payRepo, svc, cfg.DB)

	payment := r.Group("/Payment", middleware.AuthGuard())
	{
		payment.POST("/txn/order", middleware.IdempotencyMiddleware(cfg.DB, 2), controller.CreatePaymentTransactionLog)
		payment.POST("/txn/enquiry", controller.EnquiryPayment)
		payment.POST("txn/:order_number/paid", middleware.IdempotencyMiddleware(cfg.DB, 10), func(ctx *gin.Context) { controller.UpdatePaymentAndOrderStatus(ctx, models.OrderStatus.Paid) })
		payment.POST("txn/:order_number/canceled", middleware.IdempotencyMiddleware(cfg.DB, 10), func(ctx *gin.Context) { controller.UpdatePaymentAndOrderStatus(ctx, models.OrderStatus.Canceled) })
	}
}
