package models

import (
	"time"

	"gorm.io/gorm"
)

func MemberPointsMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&MemberPoint{}, &MemberPointLog{})
}

type MemberPoint struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MemberID    uint      `json:"member_id" gorm:"not null;uniqueIndex"`
	TotalPoints int       `json:"total_points" gorm:"not null;default:0"` // point * 100
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Member Member `gorm:"foreignKey:MemberID"`
}

var MemberPointLogType = struct {
	Earn   string
	Redeem string
}{Earn: "earn", Redeem: "redeem"}

type MemberPointLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MemberID  uint      `json:"member_id" gorm:"not null;index"`
	OrderID   *uint     `json:"order_id" gorm:"uniqueIndex"`
	Type      string    `json:"type" gorm:"type:varchar(20);not null"`
	Points    int       `json:"points" gorm:"not null"` // point * 100
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Member Member `gorm:"foreignKey:MemberID"`
	Order  *Order `gorm:"foreignKey:OrderID"`
}
