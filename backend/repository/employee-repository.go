package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(emp models.Employee) (models.Employee, error)
	FindOne(id int) (models.Employee, error)
	FindAll(q map[string]interface{}) ([]models.Employee, error)
	UpdateEmployee(id int, emp models.Employee) (models.Employee, error)
}

type repo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepo {
	return &repo{db}
}

func (r *repo) Create(employee models.Employee) (models.Employee, error) {

	if err := r.db.Create(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func (r *repo) FindOne(id int) (models.Employee, error) {
	var data models.Employee
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *repo) FindAll(q map[string]interface{}) ([]models.Employee, error) {
	var data []models.Employee

	err := r.db.Where(q).Find(&data).Error
	return data, err
}

func (r *repo) UpdateEmployee(id int, employee models.Employee) (models.Employee, error) {
	var data models.Employee

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(employee).Error
	return data, err
}
