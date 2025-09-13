package models

import "time"

type Order struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OrderNumber string    `gorm:"uniqueIndex;type:text" json:"order_number"`
	Status      string    `gorm:"type:text" json:"status"`
	Customer    string    `gorm:"type:text;default:guest" json:"customer"`
	Total       float64   `json:"total"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrderStatusLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	Status    string    `gorm:"type:text" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relation to Order
	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
}

type OrderMenuVariation struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	OrderID         uint    `json:"order_id"`
	MenuVariationID uint    `json:"menu_variation_id"`
	Amount          int     `json:"amount"`
	Price           float64 `json:"price"`

	// Relations
	Order         Order         `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
	MenuVariation MenuVariation `gorm:"foreignKey:MenuVariationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"menu_variation"`
}

type PaymentOrderTransactionLog struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	OrderID           uint      `json:"order_id"`
	TransactionNumber string    `gorm:"type:text" json:"transaction_number"`
	Amount            float64   `json:"amount"`
	Status            string    `gorm:"type:text" json:"status"`
	PaymentCode       string    `gorm:"type:text" json:"payment_code"`
	ExpiredAt         time.Time `json:"expired_at"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
}
