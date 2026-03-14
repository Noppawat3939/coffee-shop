package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/pagination"
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuditLogService interface {
	LogWithTx(
		tx *gorm.DB,
		employeeID *uint,
		action model.AuditLogAction,
		entity string,
		entityID uint,
		oldValue interface{},
		newValue interface{},
	) error

	FindAll(req dto.GetAuditLogRequest, p *pagination.Pagination) ([]model.AuditLog, error)
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
	action model.AuditLogAction,
	entity string,
	entityID uint,
	oldValue any,
	newValue any,
) error {
	log := model.AuditLog{
		EmployeeID: employeeID,
		Action:     action,
		Entity:     entity,
		EntityID:   entityID,
		OldData:    toJSON(oldValue),
		NewData:    toJSON(newValue),
	}

	return s.repo.Create(log, tx)
}

func (s *service) FindAll(req dto.GetAuditLogRequest, p *pagination.Pagination) ([]model.AuditLog, error) {
	filter := repository.AuditLogFilter{
		ID:        req.ID,
		Action:    req.Action,
		Entity:    req.Entity,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	return s.repo.FindAll(filter, p)
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
