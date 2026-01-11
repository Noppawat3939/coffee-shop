package routes

import (
	ctl "backend/controllers"
	"backend/middleware"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	controller := ctl.NewPaymentController(repo)

	payment := r.Group("/Payment", middleware.AuthGuard())
	{
		payment.POST("/txn/order", middleware.IdempotencyMiddleware(cfg.DB, 2), controller.CreatePaymentTransactionLog)
		payment.POST("/txn/enquiry", middleware.IdempotencyMiddleware(cfg.DB, 10), controller.EnquiryPayment)
	}
}
