package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type SessionRepo interface {
	Create(data models.Session) (models.Session, error)
	FindOne(q map[string]interface{}) (models.Session, error)
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepo {
	return &sessionRepo{db}
}

func (r *sessionRepo) Create(data models.Session) (models.Session, error) {
	if err := r.db.Create(data).Error; err != nil {
		return models.Session{}, err
	}
	return data, nil
}

func (r *sessionRepo) FindOne(q map[string]interface{}) (models.Session, error) {
	var data models.Session
	// find one not expired
	query := q
	query["expired_at > ?"] = time.Now()

	err := r.db.Where(query).First(data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
