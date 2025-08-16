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

type repo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepo {
	return &repo{db}
}

func (r *repo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *repo) FindAll() ([]models.Memu, error) {
	var data []models.Memu
	err := r.db.Preload("Variations").Find(&data).Error

	return data, err
}

func (r *repo) FindVariationAll(ids []int) ([]models.MenuVariation, error) {
	var data []models.MenuVariation

	err := r.db.Where(ids).Find(&data).Error

	return data, err
}

func (r *repo) FindOne(id int) (models.Memu, error) {
	var data models.Memu

	err := r.db.Preload("Variations").First(&data, id).Error

	return data, err
}

func (r *repo) Create(menu models.Memu, tx *gorm.DB) (models.Memu, error) {
	db := r.getDB(tx)
	if err := db.Create(&menu).Error; err != nil {
		return models.Memu{}, err
	}

	return menu, nil
}

func (r *repo) CreateMenuVariation(variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error) {
	db := r.getDB(tx)

	if err := db.Create(&variation).Error; err != nil {
		return models.MenuVariation{}, err
	}
	return variation, nil
}

func (r *repo) CreatePriceLog(priceLog models.MenuPriceLog, tx *gorm.DB) (models.MenuPriceLog, error) {
	db := r.getDB(tx)

	if err := db.Create(&priceLog).Error; err != nil {
		return models.MenuPriceLog{}, err
	}

	return priceLog, nil
}

func (r *repo) UpdateByID(id int, menu models.Memu) (models.Memu, error) {
	var data models.Memu

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(menu).Error

	return data, err
}

func (r *repo) UpdateVariationByID(id int, variation models.MenuVariation, tx *gorm.DB) (models.MenuVariation, error) {
	db := r.getDB(tx)

	var data models.MenuVariation

	if err := db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := db.Model(&data).Updates(variation).Error

	return variation, err
}
