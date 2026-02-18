package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type SessionRepo interface {
	Create(data models.Session) (models.Session, error)
	FindOne(employeeID uint) (models.Session, error)
	UpdateOne(id int) error
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepo {
	return &sessionRepo{db}
}

func (r *sessionRepo) Create(data models.Session) (models.Session, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return models.Session{}, err
	}
	return data, nil
}

func (r *sessionRepo) FindOne(employeeID uint) (models.Session, error) {
	var data models.Session

	// find one not expired
	err := r.db.Where("employee_id = ?", employeeID).Where("expired_at > ?", time.Now()).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *sessionRepo) UpdateOne(id int) error {
	return r.db.Model(&models.Session{}).Where("id = ?", id).Update("expired_at", time.Now()).Error
}
