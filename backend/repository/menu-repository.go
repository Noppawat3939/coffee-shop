package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MenuRepo interface {
	FindAll() ([]models.Memu, error)
	FindOne(id int) (models.Memu, error)
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
