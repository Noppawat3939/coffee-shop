package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitAuthRoutes(r *gin.RouterGroup) {
	repo := repository.NewEmployeeRepository(cfg.DB)
	controller := controllers.NewAuthController(repo)

	auth := r.Group("/Authen")
	{
		auth.POST("login", controller.Login)
	}
}
