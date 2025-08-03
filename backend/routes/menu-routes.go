package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialMenuRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewMenuRepository(db)
	controller := controllers.NewMenuController(repo)

	menu := r.Group("/menu")
	{
		menu.GET("/", controller.GetMenus)
		menu.GET("/:id", controller.GetMenu)
		menu.POST("/", controller.CreateMenu)
		menu.PATCH("/:id", controller.UpdateMenuByID)
	}
}
