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
	EmployeeID  uint      `json:"employee_id"`

	// Relation to OrderStatusLog
	StatusLogs          []OrderStatusLog     `gorm:"foreignKey:OrderID" json:"status_logs"`
	OrderMenuVariations []OrderMenuVariation `gorm:"foreignKey:OrderID" json:"order_menu_variations"`
	Employee            Employee             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:EmployeeID;references:ID" json:"employee"`
}

type OrderStatusLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	Status    string    `gorm:"type:text" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type OrderMenuVariation struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	OrderID         uint    `json:"order_id"`
	MenuVariationID uint    `json:"menu_variation_id"`
	Amount          int     `json:"amount"`
	Price           float64 `json:"price"`

	// Associations
	MenuVariation *MenuVariation `gorm:"foreignKey:MenuVariationID" json:"menu_variation,omitempty"`
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
