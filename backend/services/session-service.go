package services

import (
	"backend/models"
	"backend/repository"
)

type SessionService interface {
	FindOneSession(employeeID uint) (models.Session, error)
	CreateSession(data models.Session) (models.Session, error)
}

type sessionService struct {
	repo repository.SessionRepo
}

func NewSessionService(repo repository.SessionRepo) SessionService {
	return &sessionService{repo}
}

func (s *sessionService) FindOneSession(employeeID uint) (models.Session, error) {
	q := map[string]interface{}{
		"employee_id": employeeID,
	}

	v, err := s.repo.FindOne(q)
	if err != nil {
		return v, nil
	}

	return v, nil
}

func (s *sessionService) CreateSession(data models.Session) (models.Session, error) {
	v, err := s.repo.Create(data)
	if err != nil {
		return v, err
	}

	return v, nil
}
