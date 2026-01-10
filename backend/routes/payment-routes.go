package routes

import (
	ctl "backend/controllers"
	"backend/middleware"
	"backend/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	controller := ctl.NewPaymentController(repo)

	payment := r.Group("/Payment")
	{
		payment.POST("/transaction-log/order", middleware.AuthGuard(), middleware.IdempotencyMiddleware(cfg.DB, 3*time.Minute), controller.CreatePaymentTransactionLog)
		payment.POST("/transaction-log/enquiry", middleware.AuthGuard(), middleware.IdempotencyMiddleware(cfg.DB, 10*time.Minute), controller.EnquiryPayment)
	}
}
