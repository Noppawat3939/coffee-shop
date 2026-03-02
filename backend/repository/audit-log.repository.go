package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(data models.AuditLog, tx *gorm.DB) error
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db}
}

func (r *auditLogRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *auditLogRepository) Create(data models.AuditLog, tx *gorm.DB) error {
	db := r.getDB(tx)

	return db.Create(&data).Error
}
