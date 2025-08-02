package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialMenuRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewMenuRepository(db)
	ctlr := controllers.NewMenuController(repo)

	menu := r.Group("/menu")
	{
		menu.GET("/", ctlr.GetMenus)
	}
}
