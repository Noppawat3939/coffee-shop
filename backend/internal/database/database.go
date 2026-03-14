package database

import (
	"backend/config"
	md "backend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	return db, nil
}

func Migration(db *gorm.DB) error {
	fmt.Println("Running migration")

	return RunMigrations(
		db,
		md.MenuMigrate,
		md.EmployMigration,
		md.OrderMigrate,
		md.MemberMigrate,
		md.PaymenMigrate,
		md.MemberPointsMigrate,
		md.SessionMigration,
		md.AuditLogMigration,
	)
}

func RunMigrations(db *gorm.DB, migrations ...func(*gorm.DB) error) error {
	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			panic(err)
		}
	}

	return nil
}
