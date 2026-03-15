package repository

import (
	"backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepo interface {
	Create(data model.Session) error
	FindOne(employeeID uint) (model.Session, error)
	UpdateOne(id int) error
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepo {
	return &sessionRepo{db}
}

func (r *sessionRepo) Create(data model.Session) error {

	return r.db.Create(&data).Error
}

func (r *sessionRepo) FindOne(employeeID uint) (model.Session, error) {
	var data model.Session

	// find one not expired
	err := r.db.Where("employee_id = ?", employeeID).Where("expired_at > ?", time.Now()).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *sessionRepo) UpdateOne(id int) error {
	return r.db.Model(&model.Session{}).Where("id = ?", id).Update("expired_at", time.Now()).Error
}
