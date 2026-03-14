package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberRepository(db)
	pointRepo := repository.NewMemberPointRepository(db)
	memberSvc := service.NewMemberService(repo)
	pointSvc := service.NewMemberPointService(pointRepo)
	controller := controllers.NewMemberController(memberSvc, pointSvc, db)

	member := r.Group("/Members")
	{
		member.POST("/register", controller.CreateMember)
		member.POST("find", middleware.AuthGuard(), controller.GetMember)
		member.GET("", middleware.AuthGuard(), controller.GetMembers)
	}
}
