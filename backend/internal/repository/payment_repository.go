package repository

import (
	"backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	// Find all
	FindAllTransactions(q map[string]interface{}, page, limit int) ([]model.PaymentOrderTransactionLog, error)
	// Find one
	FindOneTransaction(q map[string]interface{}) (model.PaymentOrderTransactionLog, error)
	// Create
	CreatePaymentLog(data model.PaymentOrderTransactionLog, tx *gorm.DB) (model.PaymentOrderTransactionLog, error)
	// Update
	UpdatePaymentLog(q map[string]interface{}, log model.PaymentOrderTransactionLog, tx *gorm.DB) (model.PaymentOrderTransactionLog, error)
	CancelActivePaymentLog(odNumberRef string, tx *gorm.DB) error
}

type paymentRepo struct {
	db *gorm.DB
}

func (r *paymentRepo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func NewPaymentRepository(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db}
}

func (r *paymentRepo) FindOneTransaction(q map[string]interface{}) (model.PaymentOrderTransactionLog, error) {
	var log model.PaymentOrderTransactionLog

	err := r.db.Preload("Order").Where(q).First(&log).Error

	return log, err
}

func (r *paymentRepo) CreatePaymentLog(data model.PaymentOrderTransactionLog, tx *gorm.DB) (model.PaymentOrderTransactionLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return model.PaymentOrderTransactionLog{}, err
	}
	return data, nil
}

func (r *paymentRepo) UpdatePaymentLog(q map[string]interface{}, log model.PaymentOrderTransactionLog, tx *gorm.DB) (model.PaymentOrderTransactionLog, error) {
	var data model.PaymentOrderTransactionLog
	db := r.getDB(tx)

	if err := db.Where(q).First(&data).Error; err != nil {
		return data, err
	}

	if err := db.Model(&data).Updates(log).Error; err != nil {
		return data, nil
	}

	return data, nil
}

func (r *paymentRepo) CancelActivePaymentLog(odNumberRef string, tx *gorm.DB) error {
	var data model.PaymentOrderTransactionLog
	db := r.getDB(tx)

	// auto expired only status to_pay
	return db.Model(&data).Where("order_number_ref = ? AND status = ?", odNumberRef, model.OrderStatus.ToPay).Updates(map[string]interface{}{
		"status":     model.OrderStatus.Canceled,
		"expired_at": time.Now(),
	}).Error
}

func (r *paymentRepo) FindAllTransactions(q map[string]interface{}, page, limit int) ([]model.PaymentOrderTransactionLog, error) {
	var logs []model.PaymentOrderTransactionLog

	err := r.db.Joins("Order").Preload("Order.Employee").Preload("Order.Member").Where(q).Limit(limit).Offset(page).Order("id desc").Find(&logs).Error

	return logs, err
}
