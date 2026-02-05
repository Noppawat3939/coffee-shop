package db

import (
	"backend/config"
	md "backend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
	}

	fmt.Println("✅ Running migration")
	RunMigrations(db, md.MenuMigrate, md.EmployMigration, md.OrderMigrate, md.MemberMigrate, md.PaymenMigrate, md.MemberPointsMigrate)

	fmt.Println("✅ Connected to database")

	return db
}

func RunMigrations(db *gorm.DB, migrations ...func(*gorm.DB) error) {
	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			panic(err)
		}
	}

}
