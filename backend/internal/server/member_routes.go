package server

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialMemberRoutes(r *gin.RouterGroup) {
	db := cfg.DB

	repo := repository.NewMemberRepository(db)
	pointRepo := repository.NewMemberPointRepository(db)
	memberSvc := service.NewMemberService(repo)
	pointSvc := service.NewMemberPointService(pointRepo)
	handler := handler.NewMemberHandler(memberSvc, pointSvc, db)

	member := r.Group("/Members")
	{
		member.POST("/register", handler.CreateMember)
		member.POST("find", middleware.AuthGuard(), handler.GetMember)
		member.GET("", middleware.AuthGuard(), handler.GetMembers)
	}
}
