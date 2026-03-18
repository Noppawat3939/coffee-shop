package server

import (
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) registerRoutes() {
	r := s.router
	db := s.db

	r.NoRoute(func(c *gin.Context) {
		response.ErrorNotFound(c)
	})

	registerCheckRoutes(r, db)

	cfg := RouterConfig{
		Router: r,
		DB:     db,
	}

	api := r.Group("/api")

	cfg.InitAuthRoutes(api)
	cfg.InitialAuditLogRoutes(api)
	cfg.InitialMemberPointRoutes(api)
	cfg.InitialMemberRoutes(api)
	cfg.IntialEmployeeRoues(api)
	cfg.IntialMenuRoutes(api)
	cfg.IntialOrderRoutes(api)
	cfg.IntialPaymentRoutes(api)
}

func registerCheckRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database instance error"})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database down"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
