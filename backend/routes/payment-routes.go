package routes

import (
	ctl "backend/controllers"
	"backend/middleware"
	"backend/repository"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewOrderRepository(cfg.DB)
	controller := ctl.NewPaymentController(repo)

	payment := r.Group("/Payment")
	{
		payment.POST("/generate-promptpay-qr", controller.GeneratePromptPayQR)
		payment.POST("/transaction-log/order", middleware.AuthGuard(), middleware.IdempotencyMiddleware(db, 10*time.Minute), controller.CreatePaymentTransactionLog)
	}
}
