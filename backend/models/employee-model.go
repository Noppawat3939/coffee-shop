package models

import (
	"time"

	"gorm.io/gorm"
)

func EmployMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Employee{})
}

type Employee struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:text;unique;not null" json:"username"` // unique
	Password  string    `gorm:"type:text;not null" json:"-"`               // omit from JSON responses
	Name      string    `gorm:"type:text;not null" json:"name"`
	Active    bool      `gorm:"default:true" json:"active"`
	Role      string    `gorm:"type:text;not null;comment:'admin, staff'" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
