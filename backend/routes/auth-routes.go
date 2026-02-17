package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {

	repo := repository.NewEmployeeRepository(cfg.DB)
	sessionRepo := repository.NewSessionRepository(cfg.DB)
	sessionSvc := services.NewSessionService(sessionRepo)
	controller := controllers.NewAuthController(repo, sessionSvc)

	auth := r.Group("/Authen")
	{
		auth.POST("/employee/login", controller.LoginByEmployee)
		auth.POST("/employee/verification", middleware.AuthGuard(), controller.VerifyJWTByEmployee)
	}
}
