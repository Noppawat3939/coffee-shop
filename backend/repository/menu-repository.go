package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MenuRepo interface {
	FindAll() ([]models.Memu, error)
}

type repo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepo {
	return &repo{db}
}

func (r *repo) FindAll() ([]models.Memu, error) {
	var menus []models.Memu

	err := r.db.Preload("Variations").Find(&menus).Error

	return menus, err
}
