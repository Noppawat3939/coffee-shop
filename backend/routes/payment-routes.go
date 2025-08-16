package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialPaymentRoutes(r *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewPaymentController()

	payment := r.Group("/payment")
	{
		payment.POST("/generate-promptpay-qr", controller.GeneratePromptPayQR)
	}
}
