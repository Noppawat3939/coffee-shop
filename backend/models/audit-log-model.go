package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func AuditLogMigration(db *gorm.DB) error {
	return db.AutoMigrate(&AuditLog{})
}

type AuditLogAction string

var AuditAction = struct {
	Create AuditLogAction
	Update AuditLogAction
	Delete AuditLogAction
}{Create: "create", Update: "update", Delete: "delete"}

type AuditLog struct {
	ID         uint           `gorm:"primaryKey"`
	EmployeeID *uint          `gorm:"index"` // nullable (supported employee deleted)
	Action     AuditLogAction `gorm:"type:varchar(20);not null;index"`
	Entity     string         `gorm:"type:varchar(50);not null;index"`
	EntityID   uint           `gorm:"not null;index"`
	OldData    datatypes.JSON `gorm:"type:jsonb"`
	NewData    datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;index"`

	Employee *Employee `gorm:"foreignKey:EmployeeID"`
}
