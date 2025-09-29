package db

import (
	"backend/config"
	"backend/models"
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

	err = db.AutoMigrate(&models.Memu{}, &models.MenuVariation{}, &models.MenuPriceLog{}, &models.Member{}, &models.Order{}, &models.OrderStatusLog{}, &models.OrderMenuVariation{}, &models.PaymentOrderTransactionLog{})

	if err != nil {
		log.Fatalf("Failed to auto migrate tables: %v\n", err)
	}

	fmt.Println("✅ Connected to database")

	return db
}
