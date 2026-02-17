package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/pkg/types"
	"backend/repository"
	"backend/services"
	"backend/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type authController struct {
	repo       repository.EmployeeRepo
	sessionSvc services.SessionService
}

func NewAuthController(repo repository.EmployeeRepo, sessionSvc services.SessionService) *authController {
	return &authController{repo, sessionSvc}
}

func (s *authController) LoginByEmployee(c *gin.Context) {
	var req dto.LoginEmployeeRequest
	var jwt string = ""

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

	// check session not expired
	session, err := s.sessionSvc.FindOneSession(emp.ID)
	if err == nil {
		jwt = session.Value
	}

	// not session should be gen new jwt
	if jwt == "" {
		exp := time.Now().Add(time.Duration(24) * time.Hour)

		value, _ := util.GenerateJWT(emp.ID, emp.Username, exp)
		s.sessionSvc.CreateSession(models.Session{EmployeeID: &emp.ID, Value: value, Employee: &emp, ExpiredAt: exp})

		jwt = value
	}

	data := make(types.Filter)
	data["access_token"] = jwt

	util.Success(c, data)
}

func (s *authController) VerifyJWTByEmployee(c *gin.Context) {

	data, ok := util.GetUserFromHeader(c)

	if !ok {

		util.ErrorUnauthorized(c)
		return

	}

	util.Success(c, data)
}
