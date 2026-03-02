package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type AuditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db}
}

func (r *AuditLogRepository) Create(data models.AuditLog) error {
	return r.db.Create(&data).Error
}
