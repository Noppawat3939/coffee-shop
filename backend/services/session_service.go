package services

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/jwt"
	"time"
)

type SessionService interface {
	FindOneSession(employeeID uint) (model.Session, bool)
	CreateSession(data model.Session) (model.Session, error)
	ExpiredByEmployeeID(id uint) error
	GetJWT(employee model.Employee) string
}

type sessionService struct {
	repo repository.SessionRepo
}

func NewSessionService(repo repository.SessionRepo) SessionService {
	return &sessionService{repo}
}

func (s *sessionService) FindOneSession(employeeID uint) (model.Session, bool) {
	v, err := s.repo.FindOne(employeeID)

	if err != nil {
		return v, false
	}

	return v, true
}

func (s *sessionService) CreateSession(data model.Session) (model.Session, error) {
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

func (s *sessionService) GetJWT(employee model.Employee) string {
	var jwtStr string = ""

	// find session not expired
	session, found := s.FindOneSession(employee.ID)
	if found {
		jwtStr = session.Value
	}

	// not found session then gen new jwt
	if !found {
		exp := time.Now().Add(time.Duration(24) * time.Hour) // 1d
		value, _ := jwt.GenerateJWT(employee.ID, employee.Username, exp)

		data := model.Session{EmployeeID: &employee.ID, Value: value, ExpiredAt: exp, Employee: &employee}
		s.CreateSession(data)

		jwtStr = value
	}

	return jwtStr
}
