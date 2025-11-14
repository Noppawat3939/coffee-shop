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

func (s *authController) Login(c *gin.Context) {
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
