package handler

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SessionKey = "session"
)

type authHandler struct {
	repo    repository.EmployeeRepo
	authSvc service.AuthService
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func NewAuthHandler(repo repository.EmployeeRepo, authSvc service.AuthService) *authHandler {
	return &authHandler{repo, authSvc}
}

func (h *authHandler) EmployeeLoginV1(c *gin.Context) {
	response.Error(c, 406, "version not supported")
}

func (h *authHandler) LoginV2(c *gin.Context) {
	var req dto.LoginEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()

	result, err := h.authSvc.Login(req.Username, req.Password, ua, ip)
	if err != nil {
		response.ErrorUnauthorized(c)
		return
	}

	// set refresh cookie
	setCookie(c, result.RefreshToken)

	response.Success(c, AuthResponse{AccessToken: result.AccessToken})
}

func (h *authHandler) RefreshV2(c *gin.Context) {
	refresh, err := c.Cookie(SessionKey)
	if err != nil {
		response.ErrorUnauthorized(c)
		return
	}

	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()

	result, err := h.authSvc.RefreshToken(refresh, ua, ip)
	if err != nil {
		response.ErrorUnauthorized(c)
		return
	}

	setCookie(c, result.RefreshToken)

	response.Success(c, AuthResponse{AccessToken: result.AccessToken})
}

func (h *authHandler) VerifyJWTByEmployee(c *gin.Context) {
	data := auth.GetUserFromContext(c)

	response.Success(c, data)
}

func (h *authHandler) EmployeeLogout(c *gin.Context) {
	response.Error(c, 406, "version not supported")
}

func setCookie(c *gin.Context, token string) {
	maxAge := 60 * 60 * 24 * 7
	c.SetCookie(SessionKey, token, maxAge, "/", "", true, true)
}
