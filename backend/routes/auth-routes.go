package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {
	repo := repository.NewEmployeeRepository(cfg.DB)
	controller := controllers.NewAuthController(repo)

	auth := r.Group("/Authen")
	{
		auth.POST("/employee/login", controller.LoginByEmployee)
		auth.POST("/employee/verification", middleware.AuthGuard(), controller.VerifyJWTByEmployee)
	}
}
