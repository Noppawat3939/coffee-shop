package migration

import "gorm.io/gorm"

type Migration struct {
	Name string
	Up   func(*gorm.DB) error
}
