package routes

import (
	ctl "backend/controllers"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialPaymentRoutes(r *gin.RouterGroup) {
	controller := ctl.NewPaymentController()

	payment := r.Group("/Payment")
	{
		payment.POST("/generate-promptpay-qr", controller.GeneratePromptPayQR)
	}
}
