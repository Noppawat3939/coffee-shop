package controllers

import (
	"backend/dto"
	"backend/repository"
	"backend/util"
	"net/http"

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
	user, exists := c.Get("user")

	if !exists {
		util.ErrorUnauthorized(c)
		return
	}

	claims, _ := user.(*util.JWTClaims)

	data := buildUserByClaims(claims)

	util.Success(c, data)
}

func buildUserByClaims(claims *util.JWTClaims) map[string]interface{} {
	data := map[string]interface{}{
		"id":       claims.EmployeeID,
		"username": claims.Username,
		"exp":      claims.ExpiresAt,
	}

	return data
}
