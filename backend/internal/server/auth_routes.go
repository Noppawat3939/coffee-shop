package server

import (
	"backend/internal/handler"
	mdw "backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewEmployeeRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	sessionSvc := service.NewAuthService(sessionRepo, employeeRepo)
	handler := handler.NewAuthHandler(repo, sessionSvc)

	auth := r.Group("/Auth")

	v1 := auth.Group("/v1")
	{
		v1.POST("/employee/login", handler.EmployeeLoginV1)
		v1.POST("/employee/logout", mdw.AuthGuard(), handler.EmployeeLogout)
		v1.POST("/employee/verification", mdw.AuthGuard(), handler.VerifyJWTByEmployee)
	}

	v2 := auth.Group("/v2")
	{
		v2.POST("/employee/login", mdw.RateLimiter(rate.Every(time.Minute/5), 5), handler.LoginV2)
		v2.POST("/employee/refresh", handler.RefreshV2)
	}
}
