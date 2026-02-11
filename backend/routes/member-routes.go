package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberRepository(db)
	pointRepo := repository.NewMemberPointRepository(db)
	memberSvc := services.NewMemberService(repo)
	pointSvc := services.NewMemberPointService(pointRepo)
	controller := controllers.NewMemberController(memberSvc, pointSvc, db)

	member := r.Group("/Members")
	{
		member.POST("/register", controller.CreateMember)
		member.POST("find", middleware.AuthGuard(), controller.GetMember)
	}
}
