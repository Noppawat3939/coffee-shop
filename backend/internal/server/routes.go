package server

import (
	"backend/util"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	r := s.router
	db := s.db

	r.NoRoute(func(c *gin.Context) {
		util.ErrorNotFound(c)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

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
