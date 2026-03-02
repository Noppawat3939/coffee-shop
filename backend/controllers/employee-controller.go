package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/services"
	"backend/util"
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
		util.ErrorBodyInvalid(c)
		return
	}

	hash, _ := util.HashPassword(req.Password)

	employee, err := ec.repo.Create(models.Employee{
		Username: req.Username,
		Password: hash,
		Name:     req.Name,
		Role:     req.Role,
		Active:   true, // default
	})

	if err != nil {
		util.Error(c, http.StatusConflict, "failed create employee")
		return
	}

	util.Success(c, employee)
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
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, employees)
}

func (ec *employeeController) FindOne(c *gin.Context) {
	id := util.ParamToInt(c, "id")

	employee, err := ec.repo.FindOne(id)
	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, employee)
}

func (ec *employeeController) UpdateEmployee(c *gin.Context) {
	var req dto.UpdateEmployeeRequest
	user := util.GetUserFromHeader(c)
	id := util.ParamToInt(c, "id")

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	data, err := ec.service.UpdateByID(id, req, user)
	if err != nil {
		util.ErrorConflict(c)
		return
	}

	util.Success(c, data)
}
