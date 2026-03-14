package controllers

import (
	"backend/internal/auth"
	"backend/internal/dto"
	"backend/models"
	"backend/pkg/password"
	"backend/pkg/response"
	"backend/pkg/util"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type employeeController struct {
	repo    repository.EmployeeRepo
	service services.EmployeeService
	db      *gorm.DB
}

var Role = struct {
	Admin string
	Staff string
}{Admin: "admin", Staff: "staff"}

func NewEmployeeController(repo repository.EmployeeRepo, service services.EmployeeService, db *gorm.DB) *employeeController {
	return &employeeController{repo, service, db}
}

func (ec *employeeController) RegisterEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	hash, _ := password.Hash(req.Password)

	employee, err := ec.repo.Create(models.Employee{
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

func (ec *employeeController) FindAll(c *gin.Context) {
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

	employees, err := ec.repo.FindAll(q)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, employees)
}

func (ec *employeeController) FindOne(c *gin.Context) {
	id := util.ToInt((c.Param("id")))

	employee, err := ec.repo.FindOne(id)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, employee)
}

func (ec *employeeController) UpdateEmployee(c *gin.Context) {
	var req dto.UpdateEmployeeRequest
	user := auth.GetUserFromContext(c)
	id := util.ToInt((c.Param("id")))

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	data, err := ec.service.UpdateByID(id, req, user)
	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, data)
}
