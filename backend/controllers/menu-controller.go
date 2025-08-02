package controllers

import (
	"backend/helpers"
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
		helpers.ErrorNotFound(ctx)

		return
	}

	ctx.JSON(http.StatusOK, menus)
}

func (c *controller) GetMenu(ctx *gin.Context) {
	menuID := ctx.Param("id")

	id := helpers.ToInt(menuID)

	menu, err := c.repo.FindOne(id)

	if err != nil {
		helpers.ErrorNotFound(ctx)

		return
	}

	ctx.JSON(http.StatusOK, menu)
}
