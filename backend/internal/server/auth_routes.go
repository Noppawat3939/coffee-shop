package server

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewEmployeeRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	sessionSvc := service.NewAuthService(sessionRepo, employeeRepo)
	handler := handler.NewAuthHandler(repo, sessionSvc)

	auth := r.Group("/Auth")

	// old version
	v1 := auth.Group("/v1")
	{
		v1.POST("/employee/login", handler.EmployeeLoginV1)
		v1.POST("/employee/logout", middleware.AuthGuard(), handler.EmployeeLogout)
		v1.POST("/employee/verification", middleware.AuthGuard(), handler.VerifyJWTByEmployee)
	}

	// latest version
	v2 := auth.Group("/v2")
	{
		v2.POST("/employee/login", handler.LoginV2)
		v2.POST("/employee/refresh", handler.RefreshV2)
	}
}
