package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	// Create repositories
	CreateOrder(order models.Order, tx *gorm.DB) (models.Order, error)
	CreateOrderStatusLog(odLog models.OrderStatusLog, tx *gorm.DB) (models.OrderStatusLog, error)
	CreateOrderMenuVariation(odVaria models.OrderMenuVariation, tx *gorm.DB) (models.OrderMenuVariation, error)
	CreatePaymentLog(paymentOdLog models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error)

	// Find all
	FindAllOrders() ([]models.Order, error)

	// Find one
	FindOneOrder(id int) (models.Order, error)
	FindOneTransaction(txn string) (models.PaymentOrderTransactionLog, error)

	// Update one
	UpdateOrderByID(id int, order models.Order) (models.Order, error)
	UpdatePaymentLogByTransaction(txn string, txLog models.PaymentOrderTransactionLog) (models.PaymentOrderTransactionLog, error)
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

func (r *orderRepo) CreateOrder(order models.Order, tx *gorm.DB) (models.Order, error) {
	db := r.getDB(tx)
	if err := db.Create(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (r *orderRepo) CreateOrderStatusLog(odLog models.OrderStatusLog, tx *gorm.DB) (models.OrderStatusLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&odLog).Error; err != nil {
		return models.OrderStatusLog{}, err
	}
	return odLog, nil
}

func (r *orderRepo) CreateOrderMenuVariation(odVaria models.OrderMenuVariation, tx *gorm.DB) (models.OrderMenuVariation, error) {
	db := r.getDB(tx)
	if err := db.Create(&odVaria).Error; err != nil {
		return models.OrderMenuVariation{}, err
	}
	return odVaria, nil
}

func (r *orderRepo) CreatePaymentLog(paymentOdLog models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&paymentOdLog).Error; err != nil {
		return models.PaymentOrderTransactionLog{}, err
	}
	return paymentOdLog, nil
}

func (r *orderRepo) FindAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error

	return orders, err
}

func (r *orderRepo) FindOneOrder(id int) (models.Order, error) {
	var order models.Order

	err := r.db.Preload("StatusLogs").First(&order, id).Error
	return order, err
}

func (r *orderRepo) FindOneTransaction(txn string) (models.PaymentOrderTransactionLog, error) {
	var txLog models.PaymentOrderTransactionLog

	err := r.db.Where("transaction_number = ?", txn).First(&txLog).Error

	return txLog, err
}

func (r *orderRepo) UpdateOrderByID(id int, order models.Order) (models.Order, error) {
	var data models.Order

	if err := r.db.First(&data, id).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&data).Updates(order).Error

	return data, err
}

func (r *orderRepo) UpdatePaymentLogByTransaction(txn string, txLog models.PaymentOrderTransactionLog) (models.PaymentOrderTransactionLog, error) {
	var data models.PaymentOrderTransactionLog

	if err := r.db.Where("transaction_number = ?", txn).First(&data).Error; err != nil {
		return data, err
	}

	err := r.db.Model(&txLog).Updates(txLog).Error

	return txLog, err
}
