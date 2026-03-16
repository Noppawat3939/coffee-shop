package repository

import (
	"backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepo interface {
	Create(data model.Session) error
	FindByRefreshTokenHash(hash string) (*model.Session, error)
	RevokeSession(id uint) error
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

func (r *sessionRepo) FindByRefreshTokenHash(hash string) (*model.Session, error) {
	var data model.Session

	err := r.db.Where("refresh_token_hash = ?", hash).First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *sessionRepo) RevokeSession(id uint) error {
	now := time.Now()
	var data model.Session

	return r.db.Model(&data).Where("id = ?", id).Update("revoked_at", now).Error
}
