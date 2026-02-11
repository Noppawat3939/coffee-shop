package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberPointRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberPointRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	svc := services.NewMemberPointService(repo)
	memberSvc := services.NewMemberService(memberRepo)
	controller := controllers.NewMemberPointController(svc, memberSvc, db)

	point := r.Group("/Points", middleware.AuthGuard())
	{
		point.POST("/member/register/:phone_number", controller.CreateMemberPoint)
	}

}
