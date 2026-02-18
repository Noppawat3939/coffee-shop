package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/pkg/types"
	"backend/repository"
	"backend/services"
	"backend/util"
	"log"
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

func (ac *authController) EmployeeLogin(c *gin.Context) {
	var req dto.LoginEmployeeRequest
	var jwt string = ""

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	emp, err := ac.repo.FindByUsername(req.Username)
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
	session, found := ac.sessionSvc.FindOneSession(emp.ID)

	if found {
		jwt = session.Value
	}

	// not session should be gen new jwt
	if jwt == "" {
		exp := time.Now().Add(time.Duration(24) * time.Hour)
		value, _ := util.GenerateJWT(emp.ID, emp.Username, exp)
		data := models.Session{EmployeeID: &emp.ID, Value: value, ExpiredAt: exp}

		ac.sessionSvc.CreateSession(data)

		jwt = value
	}

	data := make(types.Filter)
	data["access_token"] = jwt

	util.Success(c, data)
}

func (ac *authController) VerifyJWTByEmployee(c *gin.Context) {
	data, ok := util.GetUserFromHeader(c)
	if !ok {
		util.ErrorUnauthorized(c)
		return

	}

	util.Success(c, data)
}

func (ac *authController) EmployeeLogout(c *gin.Context) {
	var msg string = ""

	data, _ := util.GetUserFromHeader(c)

	if data.ID != 0 {
		err := ac.sessionSvc.ExpiredByEmployeeID(data.ID)
		if err != nil {
			log.Println(err.Error())
			msg = "user already logged out"
		} else {
			msg = "logged out success"
		}
	}

	res := make(map[string]interface{})
	res["message"] = msg

	util.Success(c, res)
}
