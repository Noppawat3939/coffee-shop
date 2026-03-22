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
	RefreshToken(refresh, userAgent, ip string) (*RefreshResult, error)
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

type RefreshResult struct {
	AccessToken  string
	RefreshToken string
}

const (
	accessTokenExpTime  = 15 * time.Minute
	refreshTokenExpTime = 7 * 24 * time.Hour
)

func NewAuthService(sessionRepo repository.SessionRepo, employeeRepo repository.EmployeeRepo) AuthService {
	return &authService{sessionRepo, employeeRepo}
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
		employee.Role,
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

	// create new session
	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Employee:     employee,
	}, nil
}

func (s *authService) RefreshToken(refresh, userAgent, ip string) (*RefreshResult, error) {
	hash := util.HashSHA256(refresh)

	session, err := s.sessionRepo.FindByRefreshTokenHash(hash)
	if err != nil {
		return nil, appErr.ErrorUnauthorized
	}

	// check revoked or expired
	if session.RevokedAt != nil || time.Now().After(session.ExpiredAt) {
		return nil, appErr.ErrorUnauthorized
	}

	employee := session.Employee

	now := time.Now()

	// new access token
	newAccessToken, err := jwt.GenerateJWT(
		employee.ID,
		employee.Username,
		employee.Role,
		now.Add(accessTokenExpTime))

	if err != nil {
		return nil, err
	}

	// rotation refresh token
	newRefresh := uuid.NewString()
	newHash := util.HashSHA256(newRefresh)

	newSession := model.Session{
		RefreshTokenHash: newHash,
		UserAgent:        userAgent,
		IpAddress:        ip,
		EmployeeID:       &employee.ID,
		ExpiredAt:        now.Add(refreshTokenExpTime),
	}

	if err := s.sessionRepo.Create(newSession); err != nil {
		return nil, err
	}

	// revoke old token
	if err := s.sessionRepo.RevokeSession(session.ID); err != nil {
		return nil, err
	}

	return &RefreshResult{AccessToken: newAccessToken, RefreshToken: newRefresh}, nil
}
