package server

import (
	ctl "backend/internal/handler"
	repo "backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialMenuRoutes(r *gin.RouterGroup) {
	repo := repo.NewMenuRepository(cfg.DB)
	handler := ctl.NewMenuHandler(repo, cfg.DB)

	menu := r.Group("/Menus")
	{
		menu.GET("", handler.GetMenus)
		menu.GET("/:id", handler.GetMenu)
		menu.GET("/variation", handler.GetMenuVariations)
		menu.POST("/", handler.CreateMenu)
		menu.PATCH("/:id", handler.UpdateMenuByID)
		menu.PATCH("/variation/:id", handler.UpdateVariationByID)
	}
}
