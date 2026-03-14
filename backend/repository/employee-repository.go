package repository

import (
	"backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(emp model.Employee) (model.Employee, error)
	FindOne(id int) (model.Employee, error)
	FindByUsername(username string) (model.Employee, error)
	FindAll(q map[string]interface{}) ([]model.Employee, error)
	UpdateEmployeeByID(id int, emp model.Employee) (model.Employee, error)
	UpdateEmployeeByUsername(username string, emp model.Employee) (model.Employee, error)
}

type empRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepo {
	return &empRepo{db}
}

func (r *empRepo) Create(employee model.Employee) (model.Employee, error) {
	if err := r.db.Create(&employee).Error; err != nil {
		return model.Employee{}, err
	}

	return employee, nil
}

func (r *empRepo) FindOne(id int) (model.Employee, error) {
	var data model.Employee
	err := r.db.First(&data, id).Error

	return data, err
}

func (r *empRepo) FindByUsername(username string) (model.Employee, error) {
	var data model.Employee
	err := r.db.Where("username = ?", username).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, nil
	}

	return data, err
}

func (r *empRepo) FindAll(q map[string]interface{}) ([]model.Employee, error) {
	var data []model.Employee
	err := r.db.Where(q).Find(&data).Error

	return data, err
}

func (r *empRepo) UpdateEmployeeByID(id int, employee model.Employee) (model.Employee, error) {
	var data model.Employee
	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(employee).Error
	return data, err
}

func (r *empRepo) UpdateEmployeeByUsername(username string, employee model.Employee) (model.Employee, error) {
	var data model.Employee
	if err := r.db.Where("username = ?", username).First(&data).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(employee).Error
	return data, err
}
