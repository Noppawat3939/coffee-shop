package handler

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/response"
	"backend/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type menuHandler struct {
	repo repository.MenuRepo
	db   *gorm.DB
}

func NewMenuHandler(repo repository.MenuRepo, db *gorm.DB) *menuHandler {
	return &menuHandler{repo, db}
}

func (h *menuHandler) GetMenus(c *gin.Context) {
	menus, err := h.repo.FindAll()
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, menus)
}

func (h *menuHandler) GetMenuVariations(c *gin.Context) {
	idParams := c.QueryArray("id[]")

	var ids []int

	if len(idParams) != 0 {
		ids = util.StringsToInts(idParams)
	}

	data, err := h.repo.FindVariationAll(ids)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, data)
}

func (h *menuHandler) GetMenu(c *gin.Context) {
	id := util.ToInt(c.Param("id"))

	menu, err := h.repo.FindOne(id)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, menu)
}

func (h *menuHandler) CreateMenu(c *gin.Context) {
	var req dto.CreateMenuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	var data model.Memu

	err := h.db.Transaction(func(tx *gorm.DB) error {
		menu, err := h.repo.Create(model.Memu{
			Name:        req.Name,
			Description: req.Description,
			IsAvailable: req.IsAvailable,
		}, tx)

		if err != nil {
			return err
		}

		for _, v := range req.Variations {

			variation, err := h.repo.CreateMenuVariation(model.MenuVariation{
				MenuID: int(menu.ID),
				Type:   v.Type,
				Price:  v.Price,
				Image:  v.Image,
			}, tx)

			if err != nil {
				return err
			}

			_, err = h.repo.CreatePriceLog(model.MenuPriceLog{
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
		response.ErrorConflict(c)
		return
	}

	response.Success(c, data)
}

func (h *menuHandler) UpdateMenuByID(c *gin.Context) {
	id := util.ToInt(c.Param("id"))

	var body model.Memu

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "body invalid")
		return
	}

	menu, err := h.repo.UpdateByID(id, body)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, menu)
}

func (h *menuHandler) UpdateVariationByID(c *gin.Context) {
	id := util.ToInt(c.Param("id"))

	var req model.MenuVariation

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	var data model.Memu

	err := h.db.Transaction(func(tx *gorm.DB) error {
		if req.Price > 0 {
			_, err := h.repo.CreatePriceLog(model.MenuPriceLog{
				Price:           req.Price,
				MenuVariationID: uint(id),
			}, tx)

			if err != nil {
				return err
			}
		}

		_, err := h.repo.UpdateVariationByID(id, req, tx)

		if err != nil {
			return err
		}

		if err := tx.Preload("Variations").First(&data, id).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, data)
}
