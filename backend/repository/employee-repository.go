package repository

import (
	"backend/models"
	"backend/pkg/types"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(emp models.Employee) (models.Employee, error)
	FindOne(id int) (models.Employee, error)
	FindByUsername(username string) (models.Employee, error)
	FindAll(q types.Filter) ([]models.Employee, error)
	UpdateEmployeeByID(id int, emp models.Employee) (models.Employee, error)
	UpdateEmployeeByUsername(username string, emp models.Employee) (models.Employee, error)
}

type empRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepo {
	return &empRepo{db}
}

func (r *empRepo) Create(employee models.Employee) (models.Employee, error) {

	if err := r.db.Create(&employee).Error; err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func (r *empRepo) FindOne(id int) (models.Employee, error) {
	var data models.Employee
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *empRepo) FindByUsername(username string) (models.Employee, error) {
	var data models.Employee
	err := r.db.Where("username = ?", username).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, nil
	}
	return data, err
}

func (r *empRepo) FindAll(q types.Filter) ([]models.Employee, error) {
	var data []models.Employee

	err := r.db.Where(q).Find(&data).Error
	return data, err
}

func (r *empRepo) UpdateEmployeeByID(id int, employee models.Employee) (models.Employee, error) {
	var data models.Employee

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(employee).Error
	return data, err
}

func (r *empRepo) UpdateEmployeeByUsername(username string, employee models.Employee) (models.Employee, error) {
	var data models.Employee

	if err := r.db.Where("username = ?", username).First(&data).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(employee).Error
	return data, err
}
