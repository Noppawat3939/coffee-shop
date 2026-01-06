package routes

import (
	ctl "backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	repo := repository.NewOrderRepository(cfg.DB)
	controller := ctl.NewPaymentController(repo)

	payment := r.Group("/Payment")
	{
		payment.POST("/generate-promptpay-qr", controller.GeneratePromptPayQR)
		payment.POST("/transaction-log/order", controller.CreatePaymentTransactionLog)
	}
}
