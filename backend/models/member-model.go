package models

import (
	"time"

	"gorm.io/gorm"
)

func MemberMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Member{})
}

type Member struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PhoneNumber string    `json:"phone_number" gorm:"uniqueIndex:idx_phone_number"`
	Provider    string    `json:"provider"`
	FullName    string    `json:"full_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
