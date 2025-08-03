package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MenuRepo interface {
	FindAll() ([]models.Memu, error)
	FindOne(id int) (models.Memu, error)
	UpdateByID(id int, menu models.Memu) (models.Memu, error)
	Create(models.Memu) (models.Memu, error)
	CreatePriceLog(models.MenuPriceLog) (models.MenuPriceLog, error)
}

type repo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepo {
	return &repo{db}
}

func (r *repo) FindAll() ([]models.Memu, error) {
	var data []models.Memu

	err := r.db.Preload("Variations").Find(&data).Error

	return data, err
}
func (r *repo) FindOne(id int) (models.Memu, error) {
	var data models.Memu

	err := r.db.Preload("Variations").First(&data, id).Error

	return data, err
}

func (r *repo) Create(models.Memu) (models.Memu, error) {
	var data models.Memu
	return data, nil
}

func (r *repo) CreatePriceLog(models.MenuPriceLog) (models.MenuPriceLog, error) {
	var data models.MenuPriceLog

	return data, nil
}

func (r *repo) UpdateByID(id int, menu models.Memu) (models.Memu, error) {
	var data models.Memu

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(menu).Error

	return data, err
}
