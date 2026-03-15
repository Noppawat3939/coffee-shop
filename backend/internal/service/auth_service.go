package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	appErr "backend/pkg/errors"
	"backend/pkg/jwt"
	"backend/pkg/password"
	"backend/pkg/util"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(username, rawPassword, userAgent, ip string) (*LoginResult, error)
	FindOneSession(employeeID uint) (model.Session, bool)
	ExpiredByEmployeeID(id uint) error
}

type authService struct {
	sessionRepo  repository.SessionRepo
	employeeRepo repository.EmployeeRepo
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
	Employee     model.Employee
}

const (
	accessTokenExpTime  = 15 * time.Minute
	refreshTokenExpTime = 7 * 24 * time.Hour
)

func NewAuthService(sessionRepo repository.SessionRepo, employeeRepo repository.EmployeeRepo) AuthService {
	return &authService{sessionRepo, employeeRepo}
}

func (s *authService) FindOneSession(employeeID uint) (model.Session, bool) {
	v, err := s.sessionRepo.FindOne(employeeID)

	if err != nil {
		return v, false
	}

	return v, true
}

func (s *authService) ExpiredByEmployeeID(id uint) error {
	session, err := s.sessionRepo.FindOne(id)

	if err != nil {
		return err
	}

	s.sessionRepo.UpdateOne(int(session.ID))
	return nil
}

func (s *authService) Login(username, rawPassword, userAgent, ip string) (*LoginResult, error) {
	employee, err := s.employeeRepo.FindByUsername(username)

	// user not found
	if err != nil {
		password.DummyCheck()
		return nil, appErr.ErrInvalidCredential
	}

	// check hash with password
	if !password.CheckHash(rawPassword, employee.Password) {
		return nil, appErr.ErrInvalidCredential
	}

	now := time.Now()
	// generate access token
	accessExp := now.Add(accessTokenExpTime)

	accessToken, err := jwt.GenerateJWT(
		employee.ID,
		username,
		accessExp,
	)

	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken := uuid.NewString()
	hash := util.HashSHA256(refreshToken)
	exp := now.Add(refreshTokenExpTime)

	session := model.Session{
		RefreshTokenHash: hash,
		EmployeeID:       &employee.ID,
		UserAgent:        userAgent,
		IpAddress:        ip,
		ExpiredAt:        exp,
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Employee:     employee,
	}, nil
}
