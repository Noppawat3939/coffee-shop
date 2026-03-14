package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type OrderRepo interface {
	// Create repositories
	CreateOrder(order *model.Order, tx *gorm.DB) (model.Order, error)
	CreateOrderStatusLog(odLog model.OrderStatusLog, tx *gorm.DB) (model.OrderStatusLog, error)
	CreateOrderMenuVariation(odVaria model.OrderMenuVariation, tx *gorm.DB) (model.OrderMenuVariation, error)
	// Find all
	FindAllOrders(q map[string]interface{}, page, limit int) ([]model.Order, error)
	// Find one
	FindOneOrder(id int) (model.Order, error)
	FindOneOrderByOrderNumber(odNo string) (model.Order, error)
	FindOneMenuVariation(id int) (model.MenuVariation, error)
	// Update one
	UpdateOrder(q map[string]interface{}, order model.Order, tx *gorm.DB) (model.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (r *orderRepo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *orderRepo) CreateOrder(order *model.Order, tx *gorm.DB) (model.Order, error) {
	db := r.getDB(tx)
	if err := db.Create(order).Error; err != nil {
		return model.Order{}, err
	}
	return *order, nil
}

func (r *orderRepo) CreateOrderStatusLog(odLog model.OrderStatusLog, tx *gorm.DB) (model.OrderStatusLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&odLog).Error; err != nil {
		return model.OrderStatusLog{}, err
	}
	return odLog, nil
}

func (r *orderRepo) CreateOrderMenuVariation(odVaria model.OrderMenuVariation, tx *gorm.DB) (model.OrderMenuVariation, error) {
	db := r.getDB(tx)
	if err := db.Create(&odVaria).Error; err != nil {
		return model.OrderMenuVariation{}, err
	}
	return odVaria, nil
}

func (r *orderRepo) FindAllOrders(q map[string]interface{}, page, limit int) ([]model.Order, error) {
	var orders []model.Order

	err := r.db.Preload("Employee").Preload("Member").Limit(limit).Offset(page).Order("id desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepo) FindOneOrder(id int) (model.Order, error) {
	var order model.Order

	err := r.db.Preload("StatusLogs").Preload("OrderMenuVariations.MenuVariation.Menu").First(&order, id).Error
	return order, err
}

func (r *orderRepo) FindOneOrderByOrderNumber(odNo string) (model.Order, error) {
	var order model.Order

	err := r.db.Preload("Employee").Preload("StatusLogs").Preload("OrderMenuVariations.MenuVariation.Menu").Where("order_number = ?", odNo).First(&order).Error
	return order, err
}

func (r *orderRepo) FindOneMenuVariation(id int) (model.MenuVariation, error) {
	var menuVariation model.MenuVariation

	err := r.db.First(&menuVariation, id).Error
	return menuVariation, err
}

func (r *orderRepo) UpdateOrder(q map[string]interface{}, order model.Order, tx *gorm.DB) (model.Order, error) {
	var data model.Order
	db := r.getDB(tx)

	if err := db.Where(q).First(&data).Error; err != nil {
		return data, err
	}

	if err := db.Model(&data).Updates(order).Error; err != nil {
		return data, nil
	}

	return data, nil
}
