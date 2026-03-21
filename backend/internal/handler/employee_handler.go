package handler

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"

	"backend/internal/repository"
	"backend/pkg/password"
	"backend/pkg/response"
	"backend/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type employeeHandler struct {
	repo    repository.EmployeeRepo
	service service.EmployeeService
	db      *gorm.DB
}

var Role = struct {
	Admin string
	Staff string
}{Admin: "admin", Staff: "staff"}

func NewEmployeeHandler(repo repository.EmployeeRepo, service service.EmployeeService, db *gorm.DB) *employeeHandler {
	return &employeeHandler{repo, service, db}
}

func (h *employeeHandler) RegisterEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	hash := password.Hash(req.Password)

	employee, err := h.repo.Create(model.Employee{
		Username: req.Username,
		Password: hash,
		Name:     req.Name,
		Role:     req.Role,
		Active:   true, // default
	})

	if err != nil {
		response.Error(c, http.StatusConflict, "failed create employee")
		return
	}

	response.Success(c, employee)
}

func (h *employeeHandler) FindAll(c *gin.Context) {
	q := make(map[string]interface{})

	if username := c.Query("username"); username != "" {
		q["username"] = username
	}

	if name := c.Query("name"); name != "" {
		q["name"] = name
	}

	if id := c.Query("id"); id != "" {
		q["id"] = util.ToInt(id)
	}

	employees, err := h.repo.FindAll(q)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, employees)
}

func (h *employeeHandler) FindOne(c *gin.Context) {
	id := util.ToInt((c.Param("id")))

	employee, err := h.repo.FindOne(id)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, employee)
}

func (h *employeeHandler) UpdateEmployee(c *gin.Context) {
	var req dto.UpdateEmployeeRequest
	user := auth.GetUserFromContext(c)
	id := util.ToInt((c.Param("id")))

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	data, err := h.service.UpdateByID(id, req, user)
	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, data)
}
