package models

import (
	"time"

	"gorm.io/gorm"
)

// Define database schemas
type Memu struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	IsAvailable bool            `json:"is_available"`
	Variations  []MenuVariation `json:"variations,omitempty" gorm:"foreignKey:MenuID"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"deleted_at,omitempty" gorm:"index"`
}

type MenuVariation struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	MenuID        int            `json:"menu_id"`
	Type          string         `json:"type"`
	Price         float64        `json:"price"`
	Image         string         `json:"image,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	MenuPriceLogs []MenuPriceLog `json:"menu_price_logs,omitempty" gorm:"foreignKey:MenuVariationID"`

	Menu Memu `json:"menu" gorm:"foreignKey:MenuID;references:ID"`

	OrderMenuVariations []OrderMenuVariation `gorm:"foreignKey:MenuVariationID" json:"order_menu_variations,omitempty"`
}

type MenuPriceLog struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	MenuVariationID uint      `json:"menu_variation_id" gorm:"index;not null;constraint:OnUpdate:CASCADE;"`
	Price           float64   `json:"price"`
	CreatedAt       time.Time `json:"created_at"`
}
