package services

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/password"
)

type EmployeeService interface {
	UpdateByID(id int, req dto.UpdateEmployeeRequest, user *model.UserJwyToken) (model.Employee, error)
}

type employeeService struct {
	repo     repository.EmployeeRepo
	auditSvc AuditLogService
}

func NewEmployeeService(repo repository.EmployeeRepo, auditSvc AuditLogService) EmployeeService {
	return &employeeService{repo, auditSvc}
}

func (s *employeeService) UpdateByID(id int, req dto.UpdateEmployeeRequest, user *model.UserJwyToken) (model.Employee, error) {
	var oldData model.Employee
	employee, err := s.repo.FindOne(id)

	if err != nil {
		return model.Employee{}, err
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
		hash, _ := password.Hash(*req.Password)
		employee.Password = hash
	}

	result, updateErr := s.repo.UpdateEmployeeByID(id, employee)

	s.auditSvc.LogWithTx(nil, &user.ID, model.AuditAction.Update, "employee", employee.ID, oldData, employee)

	return result, updateErr
}
