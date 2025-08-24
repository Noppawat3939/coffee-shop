package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MenuRepo interface {
	FindAll() ([]models.Memu, error)
	FindOne(id int) (models.Memu, error)

	FindVariationAll(ids []int) ([]models.MenuVariation, error)

	Create(menu models.Memu, tx *gorm.DB) (models.Memu, error)
	CreatePriceLog(priceLog models.MenuPriceLog, tx *gorm.DB) (models.MenuPriceLog, error)
	CreateMenuVariation(variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error)

	UpdateByID(id int, menu models.Memu) (models.Memu, error)
	UpdateVariationByID(id int, variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error)
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepo {
	return &menuRepo{db}
}

func (r *menuRepo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *menuRepo) FindAll() ([]models.Memu, error) {
	var data []models.Memu
	err := r.db.Preload("Variations").Find(&data).Error

	return data, err
}

func (r *menuRepo) FindVariationAll(ids []int) ([]models.MenuVariation, error) {
	var data []models.MenuVariation
	query := r.db.Preload("Menu")

	if len(ids) > 0 {
		query = query.Where("menu_variations.id IN ?", ids)
	}

	err := query.Find(&data).Error

	return data, err
}

func (r *menuRepo) FindOne(id int) (models.Memu, error) {
	var data models.Memu

	err := r.db.Preload("Variations").First(&data, id).Error

	return data, err
}

func (r *menuRepo) Create(menu models.Memu, tx *gorm.DB) (models.Memu, error) {
	db := r.getDB(tx)
	if err := db.Create(&menu).Error; err != nil {
		return models.Memu{}, err
	}

	return menu, nil
}

func (r *menuRepo) CreateMenuVariation(variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error) {
	db := r.getDB(tx)

	if err := db.Create(&variation).Error; err != nil {
		return models.MenuVariation{}, err
	}
	return variation, nil
}

func (r *menuRepo) CreatePriceLog(priceLog models.MenuPriceLog, tx *gorm.DB) (models.MenuPriceLog, error) {
	db := r.getDB(tx)

	if err := db.Create(&priceLog).Error; err != nil {
		return models.MenuPriceLog{}, err
	}

	return priceLog, nil
}

func (r *menuRepo) UpdateByID(id int, menu models.Memu) (models.Memu, error) {
	var data models.Memu

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(menu).Error

	return data, err
}

func (r *menuRepo) UpdateVariationByID(id int, variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error) {
	db := r.getDB(tx)

	var data models.MenuVariation

	if err := db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := db.Model(&data).Updates(variation).Error

	return variation, err
}
