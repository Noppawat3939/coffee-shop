package services

import (
	"backend/models"
	"backend/repository"
)

type SessionService interface {
	FindOneSession(employeeID uint) (models.Session, bool)
	CreateSession(data models.Session) (models.Session, error)
	ExpiredByEmployeeID(id uint) error
}

type sessionService struct {
	repo repository.SessionRepo
}

func NewSessionService(repo repository.SessionRepo) SessionService {
	return &sessionService{repo}
}

func (s *sessionService) FindOneSession(employeeID uint) (models.Session, bool) {
	v, err := s.repo.FindOne(employeeID)

	if err != nil {
		return v, false
	}

	return v, true
}

func (s *sessionService) CreateSession(data models.Session) (models.Session, error) {
	return s.repo.Create(data)
}

func (s *sessionService) ExpiredByEmployeeID(id uint) error {
	session, err := s.repo.FindOne(id)

	if err != nil {
		return err
	}

	s.repo.UpdateOne(int(session.ID))
	return nil
}
