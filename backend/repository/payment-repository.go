package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	FindOneTransaction(q map[string]interface{}) (models.PaymentOrderTransactionLog, error)
	CreatePaymentLog(data models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error)
	UpdatePaymentLog(q map[string]interface{}, log models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error)
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

func (r *paymentRepo) FindOneTransaction(q map[string]interface{}) (models.PaymentOrderTransactionLog, error) {
	var log models.PaymentOrderTransactionLog

	err := r.db.Preload("Order").Where(q).First(&log).Error

	return log, err
}

func (r *paymentRepo) CreatePaymentLog(data models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return models.PaymentOrderTransactionLog{}, err
	}
	return data, nil
}

// TODO > change function to UpdateActivePaymentLog
func (r *paymentRepo) UpdatePaymentLog(q map[string]interface{}, log models.PaymentOrderTransactionLog, tx *gorm.DB) (models.PaymentOrderTransactionLog, error) {
	var data models.PaymentOrderTransactionLog
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
	var data models.PaymentOrderTransactionLog
	db := r.getDB(tx)

	// auto expired only status to_pay
	return db.Model(&data).Where("order_number_ref = ? AND status = ?", odNumberRef, models.OrderStatus.ToPay).Updates(map[string]interface{}{
		"status":     models.OrderStatus.Canceled,
		"expired_at": time.Now(),
	}).Error
}
