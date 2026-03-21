package service

import (
	"backend/internal/model"
	appErr "backend/pkg/errors"
	"backend/pkg/password"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockEmployeeRepo struct {
	FindByUsernameFunc func(username string) (model.Employee, error)
}

func (m *MockEmployeeRepo) FindByUsername(username string) (model.Employee, error) {
	return m.FindByUsernameFunc(username)
}

func (m *MockEmployeeRepo) Create(emp model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
}

func (m *MockEmployeeRepo) FindOne(id int) (model.Employee, error) {
	return model.Employee{}, nil
}

func (m *MockEmployeeRepo) FindAll(q map[string]interface{}) ([]model.Employee, error) {
	return []model.Employee{}, nil
}

func (m *MockEmployeeRepo) UpdateEmployeeByID(id int, emp model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
}

func (m *MockEmployeeRepo) UpdateEmployeeByUsername(username string, emp model.Employee) (model.Employee, error) {
	return model.Employee{}, nil
}

type MockSessionRepo struct {
	CreateFunc                 func(session model.Session) error
	FindByRefreshTokenHashFunc func(hash string) (*model.Session, error)
	RevokeSessionFunc          func(id uint) error
}

// mock repository
func (m *MockSessionRepo) Create(session model.Session) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(session)
	}
	return nil
}

func (m *MockSessionRepo) FindByRefreshTokenHash(hash string) (*model.Session, error) {
	if m.FindByRefreshTokenHashFunc != nil {
		return m.FindByRefreshTokenHashFunc(hash)
	}
	return nil, nil
}

func (m *MockSessionRepo) RevokeSession(id uint) error {
	if m.RevokeSessionFunc != nil {
		return m.RevokeSessionFunc(id)
	}
	return nil
}

func TestLogin_Success(t *testing.T) {
	hashed := password.Hash("mock_password")

	mockEmployeeRepo := &MockEmployeeRepo{
		FindByUsernameFunc: func(username string) (model.Employee, error) {
			return model.Employee{
				ID:       1,
				Username: username,
				Password: hashed,
			}, nil
		},
	}

	mockSessionRepo := &MockSessionRepo{
		CreateFunc: func(session model.Session) error {
			return nil
		},
	}

	svc := NewAuthService(mockSessionRepo, mockEmployeeRepo)
	res, err := svc.Login("mock_username", "mock_password", "ua", "ip")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.AccessToken)
	assert.NotEmpty(t, res.RefreshToken)
	assert.Equal(t, "mock_username", res.Employee.Username)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockEmployeeRepo := &MockEmployeeRepo{FindByUsernameFunc: func(username string) (model.Employee, error) {
		return model.Employee{}, errors.New("not found")
	}}

	mockSessionRepo := &MockSessionRepo{}

	svc := NewAuthService(mockSessionRepo, mockEmployeeRepo)
	res, err := svc.Login("mock_admin", "1234", "ua", "ip")

	assert.Nil(t, res)
	assert.ErrorIs(t, err, appErr.ErrInvalidCredential)
}

func TestLogin_InvalidPassword(t *testing.T) {
	hashed := password.Hash("mock_password")

	mockEmployeeRepo := &MockEmployeeRepo{
		FindByUsernameFunc: func(username string) (model.Employee, error) {
			return model.Employee{
				ID:       1,
				Username: "mock_username",
				Password: hashed,
			}, nil
		},
	}

	mockSessionRepo := &MockSessionRepo{}

	svc := NewAuthService(mockSessionRepo, mockEmployeeRepo)
	res, err := svc.Login("admin", "mock_password_wrong", "ua", "ip")

	assert.Nil(t, res)
	assert.ErrorIs(t, err, appErr.ErrInvalidCredential)
}

func TestLogin_CreateSessionError(t *testing.T) {
	hashed := password.Hash("mock_password")

	mockEmployeeRepo := &MockEmployeeRepo{
		FindByUsernameFunc: func(username string) (model.Employee, error) {
			return model.Employee{
				ID:       1,
				Username: "mock_username",
				Password: hashed,
			}, nil
		},
	}

	mockSessionRepo := &MockSessionRepo{
		CreateFunc: func(session model.Session) error {
			return errors.New("mock create error")
		},
	}

	svc := NewAuthService(mockSessionRepo, mockEmployeeRepo)
	res, err := svc.Login("mock_username", "mock_password", "ua", "ip")

	assert.Nil(t, res)
	assert.Error(t, err)
}
