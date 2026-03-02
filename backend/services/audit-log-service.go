package services

import (
	"backend/models"
	"backend/repository"
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuditLogService interface {
	LogWithTx(
		tx *gorm.DB,
		employeeID *uint,
		action models.AuditLogAction,
		entity string,
		entityID uint,
		oldValue interface{},
		newValue interface{},
	) error
}

type service struct {
	repo repository.AuditLogRepository
}

func NewAuditLogService(repo repository.AuditLogRepository) AuditLogService {
	return &service{repo}
}

func (s *service) LogWithTx(
	tx *gorm.DB,
	employeeID *uint,
	action models.AuditLogAction,
	entity string,
	entityID uint,
	oldValue any,
	newValue any,
) error {
	log := models.AuditLog{
		EmployeeID: employeeID,
		Action:     action,
		Entity:     entity,
		EntityID:   entityID,
		OldData:    toJSON(oldValue),
		NewData:    toJSON(newValue),
	}

	return s.repo.Create(log, tx)
}

func toJSON(v any) datatypes.JSON {
	if v == nil {
		return nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	return datatypes.JSON(b)
}
