package database

import (
	"backend/config"
	"backend/internal/model"
	"fmt"
	"time"

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

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// connection
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Connected to database")

	return db, nil
}

func Migration(db *gorm.DB) error {
	fmt.Println("Running migration")

	return RunMigrations(
		db,
		model.MenuMigrate,
		model.EmployMigration,
		model.OrderMigrate,
		model.MemberMigrate,
		model.PaymenMigrate,
		model.MemberPointsMigrate,
		model.SessionMigration,
		model.AuditLogMigration,
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
