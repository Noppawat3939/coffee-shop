package handler

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/password"
	"backend/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	repo    repository.EmployeeRepo
	authSvc service.AuthService
}

func NewAuthHandler(repo repository.EmployeeRepo, authSvc service.AuthService) *authHandler {
	return &authHandler{repo, authSvc}
}

func (h *authHandler) EmployeeLoginV1(c *gin.Context) {
	var req dto.LoginEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	emp, err := h.repo.FindByUsername(req.Username)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	ok := password.Compare(req.Password, emp.Password)

	if !ok {
		response.Error(c, http.StatusBadRequest, "invalid username or password")
		return
	}

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
	maxAge := 60 * 60 * 24 * 7
	c.SetCookie("session", result.RefreshToken, maxAge, "/", "", true, true)

	data := make(map[string]interface{})
	data["access_token"] = result.AccessToken
	response.Success(c, data)
}

func (h *authHandler) VerifyJWTByEmployee(c *gin.Context) {
	data := auth.GetUserFromContext(c)

	response.Success(c, data)
}

func (h *authHandler) EmployeeLogout(c *gin.Context) {
	var msg string = ""

	data := auth.GetUserFromContext(c)

	if data.ID != 0 {
		err := h.authSvc.ExpiredByEmployeeID(data.ID)
		if err != nil {
			log.Println(err.Error())
			msg = "user already logged out"
		} else {
			msg = "logged out success"
		}
	}

	res := make(map[string]interface{})
	res["message"] = msg

	response.Success(c, res)
}
