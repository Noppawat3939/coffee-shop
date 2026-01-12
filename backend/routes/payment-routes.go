package routes

import (
	ctl "backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	svc := services.NewPaymentService(repo)
	controller := ctl.NewPaymentController(repo, svc)

	payment := r.Group("/Payment", middleware.AuthGuard())
	{
		payment.POST("/txn/order", middleware.IdempotencyMiddleware(cfg.DB, 2), controller.CreatePaymentTransactionLog)
		payment.POST("/txn/enquiry", middleware.IdempotencyMiddleware(cfg.DB, 10), controller.EnquiryPayment)
	}
}
