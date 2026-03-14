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
	repo       repository.EmployeeRepo
	sessionSvc service.SessionService
}

func NewAuthHandler(repo repository.EmployeeRepo, sessionSvc service.SessionService) *authHandler {
	return &authHandler{repo, sessionSvc}
}

func (h *authHandler) EmployeeLogin(c *gin.Context) {
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

	data := make(map[string]interface{})
	data["access_token"] = h.sessionSvc.GetJWT(emp)

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
		err := h.sessionSvc.ExpiredByEmployeeID(data.ID)
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
