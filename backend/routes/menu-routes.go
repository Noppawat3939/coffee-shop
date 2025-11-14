package routes

import (
	ctl "backend/controllers"
	repo "backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialMenuRoutes(r *gin.RouterGroup) {
	repo := repo.NewMenuRepository(cfg.DB)
	controller := ctl.NewMenuController(repo, cfg.DB)

	menu := r.Group("/Menus")
	{
		menu.GET("", controller.GetMenus)
		menu.GET("/:id", controller.GetMenu)
		menu.GET("/variation", controller.GetMenuVariations)
		menu.POST("/", controller.CreateMenu)
		menu.PATCH("/:id", controller.UpdateMenuByID)
		menu.PATCH("/variation/:id", controller.UpdateVariationByID)
	}
}
