package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberRoutes(r *gin.RouterGroup) {
	repo := repository.NewMemberRepository(cfg.DB)
	controller := controllers.NewMemberController(repo, cfg.DB)

	member := r.Group("/Members")
	{
		member.POST("/register", controller.CreateMember)
		member.POST("", controller.GetMember)
	}
}
