package routes

import (
	"backend/controllers"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberRepository(db)
	pointRepo := repository.NewMemberPointRepository(db)
	pointSvc := services.NewMemberPointService(pointRepo)
	controller := controllers.NewMemberController(repo, pointSvc, db)

	member := r.Group("/Members")
	{
		member.POST("/register", controller.CreateMember)
		member.POST("", controller.GetMember)
	}
}
