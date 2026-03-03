package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/util"
)

type EmployeeService interface {
	UpdateByID(id int, req dto.UpdateEmployeeRequest, user *models.UserJwyToken) (models.Employee, error)
}

type employeeService struct {
	repo     repository.EmployeeRepo
	auditSvc AuditLogService
}

func NewEmployeeService(repo repository.EmployeeRepo, auditSvc AuditLogService) EmployeeService {
	return &employeeService{repo, auditSvc}
}

func (s *employeeService) UpdateByID(id int, req dto.UpdateEmployeeRequest, user *models.UserJwyToken) (models.Employee, error) {
	var oldData models.Employee
	employee, err := s.repo.FindOne(id)

	if err != nil {
		return models.Employee{}, err
	}

	oldData = employee

	if req.Name != nil {
		employee.Name = *req.Name
	}

	if req.Username != nil {
		employee.Username = *req.Username
	}

	if req.Role != nil {
		employee.Role = *req.Role
	}

	if req.Active != nil {
		employee.Active = *req.Active
	}

	if req.Password != nil {
		hash, _ := util.HashPassword(*req.Password)
		employee.Password = hash
	}

	result, updateErr := s.repo.UpdateEmployeeByID(id, employee)

	s.auditSvc.LogWithTx(nil, &user.ID, models.AuditAction.Update, "employee", employee.ID, oldData, employee)

	return result, updateErr
}
