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

	MemberPoint *MemberPoint `json:"member_point,omitempty" gorm:"foreignKey:MemberID"`
}

// filter
type MemberFilter struct {
	PhoneNumber string
	FullName    string
}

type MemberResponse struct {
	ID          uint         `json:"id"`
	PhoneNumber string       `json:"phone_number"`
	Provider    string       `json:"provider"`
	FullName    string       `json:"full_name"`
	CreatedAt   time.Time    `json:"created_at"`
	MemberPoint *MemberPoint `json:"member_point,omitempty"`
}
