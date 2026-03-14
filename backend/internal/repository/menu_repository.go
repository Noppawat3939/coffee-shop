package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type MenuRepo interface {
	FindAll() ([]model.Memu, error)
	FindOne(id int) (model.Memu, error)

	FindVariationAll(ids []int) ([]model.MenuVariation, error)

	Create(menu model.Memu, tx *gorm.DB) (model.Memu, error)
	CreatePriceLog(priceLog model.MenuPriceLog, tx *gorm.DB) (model.MenuPriceLog, error)
	CreateMenuVariation(variation model.MenuVariation, tx *gorm.DB) (model.MenuVariation, error)

	UpdateByID(id int, menu model.Memu) (model.Memu, error)
	UpdateVariationByID(id int, variation model.MenuVariation, tx *gorm.DB) (model.MenuVariation, error)
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

func (r *menuRepo) FindAll() ([]model.Memu, error) {
	var data []model.Memu
	err := r.db.Preload("Variations").Find(&data).Error

	return data, err
}

func (r *menuRepo) FindVariationAll(ids []int) ([]model.MenuVariation, error) {
	var data []model.MenuVariation
	query := r.db.Preload("Menu")

	if len(ids) > 0 {
		query = query.Where("menu_variations.id IN ?", ids)
	}

	err := query.Find(&data).Error

	return data, err
}

func (r *menuRepo) FindOne(id int) (model.Memu, error) {
	var data model.Memu

	err := r.db.Preload("Variations").First(&data, id).Error

	return data, err
}

func (r *menuRepo) Create(menu model.Memu, tx *gorm.DB) (model.Memu, error) {
	db := r.getDB(tx)
	if err := db.Create(&menu).Error; err != nil {
		return model.Memu{}, err
	}

	return menu, nil
}

func (r *menuRepo) CreateMenuVariation(variation model.MenuVariation, tx *gorm.DB) (model.MenuVariation, error) {
	db := r.getDB(tx)

	if err := db.Create(&variation).Error; err != nil {
		return model.MenuVariation{}, err
	}
	return variation, nil
}

func (r *menuRepo) CreatePriceLog(priceLog model.MenuPriceLog, tx *gorm.DB) (model.MenuPriceLog, error) {
	db := r.getDB(tx)

	if err := db.Create(&priceLog).Error; err != nil {
		return model.MenuPriceLog{}, err
	}

	return priceLog, nil
}

func (r *menuRepo) UpdateByID(id int, menu model.Memu) (model.Memu, error) {
	var data model.Memu

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(menu).Error

	return data, err
}

func (r *menuRepo) UpdateVariationByID(id int, variation model.MenuVariation, tx *gorm.DB) (model.MenuVariation, error) {
	db := r.getDB(tx)

	var data model.MenuVariation

	if err := db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := db.Model(&data).Updates(variation).Error

	return variation, err
}
