package controllers

import (
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	repo repository.MenuRepo
}

func NewMenuController(repo repository.MenuRepo) *controller {
	return &controller{repo}
}

func (c *controller) GetMenus(ctx *gin.Context) {
	menus, err := c.repo.FindAll()
	if err != nil {
		data := map[string]interface{}{
			"success": false,
		}
		ctx.JSON(http.StatusInternalServerError, data)
	}

	ctx.JSON(http.StatusOK, menus)
}
