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

type menuController struct {
	repo repository.MenuRepo
	db   *gorm.DB
}

func NewMenuController(repo repository.MenuRepo, db *gorm.DB) *menuController {
	return &menuController{repo, db}
}

func (mc *menuController) GetMenus(c *gin.Context) {
	menus, err := mc.repo.FindAll()
	if err != nil {
		hlp.ErrorNotFound(c)

		return
	}

	hlp.Success(c, menus)
}

func (mc *menuController) GetMenuVariations(c *gin.Context) {
	idParams := c.QueryArray("id[]")

	var ids []int

	if len(idParams) != 0 {
		ids = hlp.StringsToInts(idParams)
	}

	data, err := mc.repo.FindVariationAll(ids)

	if err != nil {
		hlp.ErrorNotFound(c)
		return
	}

	hlp.Success(c, data)
}

func (mc *menuController) GetMenu(c *gin.Context) {
	id := hlp.ParamToInt(c, "id")

	menu, err := mc.repo.FindOne(id)

	if err != nil {
		hlp.ErrorNotFound(c)

		return
	}

	hlp.Success(c, menu)
}

func (mc *menuController) CreateMenu(c *gin.Context) {
	var req dto.CreateMenuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(c)
		return
	}

	var data models.Memu

	err := mc.db.Transaction(func(tx *gorm.DB) error {
		menu, err := mc.repo.Create(models.Memu{
			Name:        req.Name,
			Description: req.Description,
			IsAvailable: req.IsAvailable,
		}, tx)

		if err != nil {
			return err
		}

		for _, v := range req.Variations {

			variation, err := mc.repo.CreateMenuVariation(models.MenuVariation{
				MenuID: int(menu.ID),
				Type:   v.Type,
				Price:  v.Price,
				Image:  v.Image,
			}, tx)

			if err != nil {
				return err
			}

			_, err = mc.repo.CreatePriceLog(models.MenuPriceLog{
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
		hlp.Error(c, http.StatusConflict, "failed create menu")
		return
	}

	hlp.Success(c, data)
}

func (mc *menuController) UpdateMenuByID(c *gin.Context) {
	id := hlp.ParamToInt(c, "id")

	var body models.Memu

	if err := c.ShouldBindJSON(&body); err != nil {
		hlp.Error(c, http.StatusBadRequest, "body invalid")
		return
	}

	menu, err := mc.repo.UpdateByID(id, body)

	if err != nil {
		hlp.Error(c, http.StatusConflict, "failed update menu id %d ", id)
		return
	}

	hlp.Success(c, menu)
}

func (mc *menuController) UpdateVariationByID(c *gin.Context) {
	id := hlp.ParamToInt(c, "id")

	var req models.MenuVariation

	if err := c.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(c)
		return
	}

	var data models.Memu

	err := mc.db.Transaction(func(tx *gorm.DB) error {
		if req.Price > 0 {
			_, err := mc.repo.CreatePriceLog(models.MenuPriceLog{
				Price:           req.Price,
				MenuVariationID: uint(id),
			}, tx)

			if err != nil {
				return err
			}
		}

		_, err := mc.repo.UpdateVariationByID(id, req, tx)

		if err != nil {
			return err
		}

		if err := tx.Preload("Variations").First(&data, id).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		hlp.Error(c, http.StatusConflict, "failed update menu variation id: %d", id)
		return
	}

	hlp.Success(c, data)
}
