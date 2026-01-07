package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func PaymenMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&PaymentOrderTransactionLog{}, &IdempotencyKey{})
}

type PaymentOrderTransactionLog struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	OrderID           uint      `json:"order_id"`
	TransactionNumber string    `gorm:"type:text" json:"transaction_number"`
	Amount            float64   `json:"amount"`
	Status            string    `gorm:"type:text" json:"status"`
	PaymentCode       string    `gorm:"type:text" json:"payment_code"`
	QRSignature       string    `gorm:"type:text" json:"qr_signature"`
	ExpiredAt         time.Time `json:"expired_at"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
}

type IdempotencyKey struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Key        string
	Endpoint   string
	Response   datatypes.JSON
	StatusCode int
	ExpiredAt  time.Time
	CreatedAt  time.Time
}
