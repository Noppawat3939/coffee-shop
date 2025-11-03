package routes

import (
	"backend/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.NoRoute(
		func(c *gin.Context) { util.ErrorNotFound(c) })

	api := r.Group("/api")

	IntialMenuRoutes(api, db)
	IntialPaymentRoutes(api, db)
	InitialMemberRoutes(api, db)
	IntialOrderRoutes(api, db)
}
