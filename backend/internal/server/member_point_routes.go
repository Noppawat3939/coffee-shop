package server

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberPointRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberPointRepository(db)
	memberRepo := repository.NewMemberRepository(db)
	svc := service.NewMemberPointService(repo)
	memberSvc := service.NewMemberService(memberRepo)
	handler := handler.NewMemberPointHandler(svc, memberSvc, db)

	point := r.Group("/Points", middleware.AuthGuard())
	{
		point.POST("/member/register/:phone_number", handler.CreateMemberPoint)
	}

}
