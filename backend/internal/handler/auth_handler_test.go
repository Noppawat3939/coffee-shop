package handler

import (
	"backend/internal/dto"
	"backend/internal/service"
	"errors"

	appErr "backend/pkg/errors"
	"backend/pkg/util"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	LoginFunc        func(username, password, userAgent, ip string) (*service.LoginResult, error)
	RefreshTokenFunc func(refresh, ua, ip string) (*service.RefreshResult, error)
}

func (m *MockAuthService) Login(username, password, userAgent, ip string) (*service.LoginResult, error) {
	return m.LoginFunc(username, password, userAgent, ip)
}

func (m *MockAuthService) RefreshToken(refresh, ua, ip string) (*service.RefreshResult, error) {
	if m.RefreshTokenFunc != nil {
		return m.RefreshTokenFunc(refresh, ua, ip)
	}
	return nil, nil
}

func TestLoginV2Handler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAuthService{
		LoginFunc: func(username, password, userAgent, ip string) (*service.LoginResult, error) {
			return &service.LoginResult{
				AccessToken:  "mock_access",
				RefreshToken: "mock_refresh",
			}, nil
		},
		RefreshTokenFunc: func(refresh, ua, ip string) (*service.RefreshResult, error) {
			return &service.RefreshResult{
				AccessToken:  "mock_access",
				RefreshToken: "mock_refresh",
			}, nil
		},
	}

	h := NewAuthHandler(nil, mockSvc)
	r := gin.Default()
	path := "/api/Auth/v2/employee/login"
	r.POST(path, h.LoginV2)

	body := dto.LoginEmployeeRequest{
		Username: "admin",
		Password: "1234",
	}

	jsonBody := util.JSONParse(body)

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	t.Log("status:", w.Code)
	t.Log("body:", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	cookies := w.Result().Cookies()
	found := false
	for _, c := range cookies {
		if c.Name == SessionKey {
			found = true
		}
	}
	assert.True(t, found, "session cookie should be set")
}

func TestLoginV2Handler_InvalidCredential(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAuthService{
		LoginFunc: func(username, password, ua, ip string) (*service.LoginResult, error) {
			return nil, appErr.ErrInvalidCredential
		},
	}

	h := NewAuthHandler(nil, mockSvc)

	r := gin.Default()
	path := "/api/Auth/v2/employee/login"
	r.POST(path, h.LoginV2)

	body := `{"username":"test","password":"wrong"}`

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	t.Log("status:", w.Code)
	t.Log("body:", w.Body.String())

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLoginV2Handler_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAuthService{}

	h := NewAuthHandler(nil, mockSvc)

	r := gin.Default()
	path := "/api/Auth/v2/employee/login"
	r.POST(path, h.LoginV2)

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	t.Log("status:", w.Code)
	t.Log("body:", w.Body.String())

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginV2Handler_InternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAuthService{
		LoginFunc: func(username, password, ua, ip string) (*service.LoginResult, error) {
			return nil, errors.New("db error")
		},
	}

	h := NewAuthHandler(nil, mockSvc)

	r := gin.Default()
	path := "/api/Auth/v2/employee/login"
	r.POST(path, h.LoginV2)

	body := `{"username":"test","password":"1234"}`

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	t.Log("status:", w.Code)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
