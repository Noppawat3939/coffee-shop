package services

import (
	"backend/models"
	"backend/repository"
	"backend/util"
	"time"
)

type SessionService interface {
	FindOneSession(employeeID uint) (models.Session, bool)
	CreateSession(data models.Session) (models.Session, error)
	ExpiredByEmployeeID(id uint) error
	GetJWT(employee models.Employee) string
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

func (s *sessionService) GetJWT(employee models.Employee) string {
	var jwt string = ""

	// find session not expired
	session, found := s.FindOneSession(employee.ID)
	if found {
		jwt = session.Value
	}

	// not found session then gen new jwt
	if !found {
		exp := time.Now().Add(time.Duration(24) * time.Hour) // 1d
		value, _ := util.GenerateJWT(employee.ID, employee.Username, exp)
		data := models.Session{EmployeeID: &employee.ID, Value: value, ExpiredAt: exp, Employee: &employee}
		s.CreateSession(data)

		jwt = value
	}

	return jwt
}
