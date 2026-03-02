package services

import (
	"backend/models"
	"backend/repository"
	"encoding/json"

	"gorm.io/datatypes"
)

type AuditLogService struct {
	repo repository.AuditLogRepository
}

func NewAuditLogService(repo repository.AuditLogRepository) *AuditLogService {
	return &AuditLogService{repo}
}

func (s *AuditLogService) Log(employeeID *uint, action models.AuditLogAction, entity string, entityID uint, oldValue interface{}, newValue interface{}) {
	log := models.AuditLog{
		EmployeeID: employeeID,
		Action:     action,
		Entity:     entity,
		EntityID:   entityID,
		OldData:    toJSON(oldValue),
		NewData:    toJSON(newValue),
	}

	go s.repo.Create(log)
}

func toJSON(v interface{}) datatypes.JSON {
	if v == nil {
		return nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	return datatypes.JSON(b)
}
