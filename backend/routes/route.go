package routes

import (
	"backend/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterConfig struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.NoRoute(
		func(c *gin.Context) { util.ErrorNotFound(c) })

	cfg := RouterConfig{
		Router: r,
		DB:     db,
	}

	api := r.Group("/api")

	cfg.InitAuthRoutes(api)
	cfg.InitialMemberRoutes(api)
	cfg.IntialEmployeeRoues(api)
	cfg.IntialMenuRoutes(api)
	cfg.IntialOrderRoutes(api)
	cfg.IntialPaymentRoutes(api)
	cfg.InitialMemberPointRoutes(api)
}
