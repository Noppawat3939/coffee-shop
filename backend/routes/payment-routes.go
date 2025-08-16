package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialPaymentRoutes(r *gin.RouterGroup, db *gorm.DB) {

	r.Group("/payment")
}
