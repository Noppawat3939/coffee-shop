package controllers

import (
	"backend/dto"
	"backend/repository"
	"backend/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authController struct {
	repo repository.EmployeeRepo
}

func NewAuthController(repo repository.EmployeeRepo) *authController {
	return &authController{repo}
}

func (s *authController) LoginByEmployee(c *gin.Context) {
	var req dto.LoginEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	emp, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	ok := util.CheckPasswordHash(req.Password, emp.Password)

	if !ok {
		util.Error(c, http.StatusBadRequest, "invalid username or password")
		return
	}

	jwt, _ := util.GenerateJWT(emp.ID, emp.Username)

	data := make(map[string]interface{})
	data["access_token"] = jwt

	util.Success(c, data)
}

func (s *authController) VerifyJWTByEmployee(c *gin.Context) {
	authHeader := util.GetAuthHeader(c)

	var authPrefix = "Bearer "

	if authHeader == "" || !strings.HasPrefix(authHeader, authPrefix) {
		util.ErrorUnauthorized(c)
		return
	}

	jwt := strings.TrimPrefix(authHeader, authPrefix)

	_, err := util.ParseJWT(jwt)

	if err != nil {
		util.ErrorUnauthorized(c)
		return
	}

	util.Success(c)
}
