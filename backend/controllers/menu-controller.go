package controllers

import (
	"backend/helpers"
	"backend/models"
	"backend/repository"
	"fmt"
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

	helpers.Success(ctx, menus)
}

func (c *controller) GetMenu(ctx *gin.Context) {

	id := helpers.ParamToInt(ctx, "id")

	menu, err := c.repo.FindOne(id)

	if err != nil {
		helpers.ErrorNotFound(ctx)

		return
	}

	helpers.Success(ctx, menu)
}

func (c *controller) CreateMenu(ctx *gin.Context) {
	var body models.Memu

	if err := ctx.ShouldBindJSON(&body); err != nil {
		helpers.Error(ctx, http.StatusBadRequest, "body invalid")
		return
	}

	menu, err := c.repo.Create(body)

	if err != nil {
		helpers.Error(ctx, http.StatusConflict, "failed create menu")
		return
	}

	helpers.Success(ctx, menu)
}

func (c *controller) UpdateMenuByID(ctx *gin.Context) {
	id := helpers.ParamToInt(ctx, "id")

	var body models.Memu

	if err := ctx.ShouldBindJSON(&body); err != nil {
		helpers.Error(ctx, http.StatusBadRequest, "body invalid")
		return
	}

	fmt.Println(id)

	menu, err := c.repo.UpdateByID(id, body)

	if err != nil {
		helpers.Error(ctx, http.StatusConflict, "failed update menu id %d ", id)
		return
	}

	helpers.Success(ctx, menu)
}
