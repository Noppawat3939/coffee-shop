package server

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {

	repo := repository.NewEmployeeRepository(cfg.DB)
	sessionRepo := repository.NewSessionRepository(cfg.DB)
	sessionSvc := service.NewSessionService(sessionRepo)
	handler := handler.NewAuthHandler(repo, sessionSvc)

	auth := r.Group("/Authen")
	{
		auth.POST("/employee/login", handler.EmployeeLogin)
		auth.POST("/employee/logout", middleware.AuthGuard(), handler.EmployeeLogout)
		auth.POST("/employee/verification", middleware.AuthGuard(), handler.VerifyJWTByEmployee)
	}
}
