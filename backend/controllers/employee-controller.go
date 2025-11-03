package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type employeeController struct {
	repo repository.EmployeeRepo
	db   *gorm.DB
}

var Role = struct {
	Admin string
	Staff string
}{Admin: "admin", Staff: "staff"}

func NewEmployeeController(repo repository.EmployeeRepo, db *gorm.DB) *employeeController {
	return &employeeController{repo, db}
}

func (ec *employeeController) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	employee, err := ec.repo.Create(models.Employee{
		Username: req.Username,
		Password: req.Password, // hash
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
	filter := make(map[string]interface{})

	if username := c.Query("username"); username != "" {
		filter["username"] = username
	}

	if name := c.Query("name"); name != "" {
		filter["name"] = name
	}

	if id := c.Query("id"); id != "" {
		filter["id"] = util.ToInt(id)
	}

	employees, err := ec.repo.FindAll(filter)
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

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	// TODO: handle update by id

	util.Success(c, nil)
}
