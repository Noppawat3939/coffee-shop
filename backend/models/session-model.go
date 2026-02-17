package models

import (
	"time"

	"gorm.io/gorm"
)

func SessionMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Session{})
}

type Session struct {
	ID         uint      `gorm:"primaryKey;autoIncreasement" json:"id"`
	Value      string    `json:"value"`
	EmployeeID *uint     `json:"employee_id"`
	ExpiredAt  time.Time `json:"expired_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`

	// relation
	Employee *Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:EmployeeID;references:ID" json:"employee"`
}
