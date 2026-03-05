package repository

import (
	"backend/models"
	"backend/util"
	"time"

	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(data models.AuditLog, tx *gorm.DB) error
	FindAll(filter AuditLogFilter, p *util.Pagination) ([]models.AuditLog, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

type AuditLogFilter struct {
	ID        *uint
	Action    *string
	Entity    *string
	StartDate *time.Time
	EndDate   *time.Time
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

func (r *auditLogRepository) FindAll(filter AuditLogFilter, p *util.Pagination) ([]models.AuditLog, error) {
	var data []models.AuditLog

	query := r.db.Model(&models.AuditLog{}).Preload("Employee").Scopes(buildFilter(filter))

	query = p.Apply(query)
	err := query.Find(&data).Error

	return data, err
}

func buildFilter(filter AuditLogFilter) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != nil {
			db = db.Where("id = ?", *filter.ID)
		}
		if filter.Action != nil {
			db = db.Where("action = ?", *filter.Action)
		}
		if filter.Entity != nil {
			db = db.Where("entity = ?", *filter.EndDate)
		}
		if filter.StartDate != nil && filter.EndDate != nil {
			db = db.Where("created_at BETWEEN ? AND ?", *filter.StartDate, *filter.EndDate)
		}

		return db
	}
}
