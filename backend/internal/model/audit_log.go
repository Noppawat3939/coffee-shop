package model

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
	ID         uint           `gorm:"primaryKey" json:"id"`
	EmployeeID *uint          `gorm:"index" json:"employee_id"` // nullable (supported employee deleted)
	Action     AuditLogAction `gorm:"type:varchar(20);not null;index" json:"action"`
	Entity     string         `gorm:"type:varchar(50);not null;index" json:"entity"`
	EntityID   uint           `gorm:"not null;index" json:"entity_id"`
	OldData    datatypes.JSON `gorm:"type:jsonb" json:"old_data"`
	NewData    datatypes.JSON `gorm:"type:jsonb" json:"new_data"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;index" json:"created_at"`

	Employee *Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
}
