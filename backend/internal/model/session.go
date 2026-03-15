package model

import (
	"time"

	"gorm.io/gorm"
)

func SessionMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Session{})
}

type Session struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	RefreshTokenHash string `gorm:"index:not null" json:"refresh_token_hash"`
	UserAgent        string `json:"user_agent"`
	IpAddress        string `json:"ip_address"`
	EmployeeID       *uint  `gorm:"index" json:"employee_id"`

	ExpiredAt time.Time  `gorm:"not null" json:"expired_at"`
	RevokedAt *time.Time `json:"revoked_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// relation
	Employee *Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:EmployeeID;references:ID" json:"employee"`
}
