package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberPointRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberPointRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	svc := service.NewMemberPointService(repo)
	memberSvc := service.NewMemberService(memberRepo)
	controller := controllers.NewMemberPointController(svc, memberSvc, db)

	point := r.Group("/Points", middleware.AuthGuard())
	{
		point.POST("/member/register/:phone_number", controller.CreateMemberPoint)
	}

}
