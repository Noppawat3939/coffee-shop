package routes

import (
	hlp "backend/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.NoRoute(
		func(c *gin.Context) { hlp.ErrorNotFound(c) })

	api := r.Group("/api")

	IntialMenuRoutes(api, db)
	IntialPaymentRoutes(api, db)
	InitialMemberRoutes(api, db)
}
