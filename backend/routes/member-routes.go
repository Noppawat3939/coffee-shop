package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitialMemberRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewMemberRepository(db)
	controller := controllers.NewMemberController(repo, db)

	member := r.Group("/members")
	{
		member.POST("/register", controller.CreateMember)
		member.POST("", controller.GetMember)
	}
}
