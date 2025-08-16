package controllers

import (
	"backend/dto"
	hlp "backend/helpers"
	"backend/models"
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	repo repository.MenuRepo
	db   *gorm.DB
}

func NewMenuController(repo repository.MenuRepo, db *gorm.DB) *controller {
	return &controller{repo, db}
}

func (c *controller) GetMenus(ctx *gin.Context) {
	menus, err := c.repo.FindAll()
	if err != nil {
		hlp.ErrorNotFound(ctx)

		return
	}

	hlp.Success(ctx, menus)
}

func (c *controller) GetMenuVariations(ctx *gin.Context) {
	id := ctx.Query("id")

	var ids []int

	if id != "" {
		ids = hlp.IdStringToInts(id, ",")
	}

	data, err := c.repo.FindVariationAll(ids)

	if err != nil {
		hlp.ErrorNotFound(ctx)
		return
	}

	hlp.Success(ctx, data)
}

func (c *controller) GetMenu(ctx *gin.Context) {

	id := hlp.ParamToInt(ctx, "id")

	menu, err := c.repo.FindOne(id)

	if err != nil {
		hlp.ErrorNotFound(ctx)

		return
	}

	hlp.Success(ctx, menu)
}

func (c *controller) CreateMenu(ctx *gin.Context) {
	var req dto.CreateMenuRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(ctx)
		return
	}

	var data models.Memu

	err := c.db.Transaction(func(tx *gorm.DB) error {
		menu, err := c.repo.Create(models.Memu{
			Name:        req.Name,
			Description: req.Description,
			IsAvailable: req.IsAvailable,
		}, tx)

		if err != nil {
			return err
		}

		for _, v := range req.Variations {

			variation, err := c.repo.CreateMenuVariation(models.MenuVariation{
				MenuID: int(menu.ID),
				Type:   v.Type,
				Price:  v.Price,
				Image:  v.Image,
			}, tx)

			if err != nil {
				return err
			}

			_, err = c.repo.CreatePriceLog(models.MenuPriceLog{
				MenuVariationID: variation.ID,
				Price:           variation.Price,
			}, tx)

			if err != nil {
				return err
			}
		}

		if err := tx.Preload("Variations.MenuPriceLogs").First(&data, menu.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		hlp.Error(ctx, http.StatusConflict, "failed create menu")
		return
	}

	hlp.Success(ctx, data)
}

func (c *controller) UpdateMenuByID(ctx *gin.Context) {
	id := hlp.ParamToInt(ctx, "id")

	var body models.Memu

	if err := ctx.ShouldBindJSON(&body); err != nil {
		hlp.Error(ctx, http.StatusBadRequest, "body invalid")
		return
	}

	menu, err := c.repo.UpdateByID(id, body)

	if err != nil {
		hlp.Error(ctx, http.StatusConflict, "failed update menu id %d ", id)
		return
	}

	hlp.Success(ctx, menu)
}

func (c *controller) UpdateVariationByID(ctx *gin.Context) {
	id := hlp.ParamToInt(ctx, "id")

	var req models.MenuVariation

	if err := ctx.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(ctx)
		return
	}

	var data models.Memu

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if req.Price > 0 {
			_, err := c.repo.CreatePriceLog(models.MenuPriceLog{
				Price:           req.Price,
				MenuVariationID: uint(id),
			}, tx)

			if err != nil {
				return err
			}
		}

		_, err := c.repo.UpdateVariationByID(id, req, tx)

		if err != nil {
			return err
		}

		if err := tx.Preload("Variations").First(&data, id).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		hlp.Error(ctx, http.StatusConflict, "failed update menu variation id: %d", id)
		return
	}

	hlp.Success(ctx, data)
}
