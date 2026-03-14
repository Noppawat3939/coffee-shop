package controllers

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/pkg/password"
	"backend/pkg/response"
	"backend/services"
	"log"
	"net/http"

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

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	emp, err := ac.repo.FindByUsername(req.Username)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	ok := password.Compare(req.Password, emp.Password)

	if !ok {
		response.Error(c, http.StatusBadRequest, "invalid username or password")
		return
	}

	data := make(map[string]interface{})
	data["access_token"] = ac.sessionSvc.GetJWT(emp)

	response.Success(c, data)
}

func (ac *authController) VerifyJWTByEmployee(c *gin.Context) {
	data := auth.GetUserFromContext(c)

	response.Success(c, data)
}

func (ac *authController) EmployeeLogout(c *gin.Context) {
	var msg string = ""

	data := auth.GetUserFromContext(c)

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

	response.Success(c, res)
}
