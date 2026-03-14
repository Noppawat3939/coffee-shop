package main

import (
	"backend/config"

	"backend/internal/database"
	"backend/internal/migration"
	"log"
	"os"

	"gorm.io/gorm"
)

// How to run:
// - make migration name={name}
func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// no agr show usage
	if len(os.Args) < 2 {
		log.Fatal("Please provide migration name, ex: go run cmd/migration/main.go {file name}")
	}

	cmd := os.Args[1]

	var selected *migration.Migration

	for _, m := range migration.AllMigrations {
		if m.Name == cmd {
			selected = &m
			break
		}
	}

	if selected == nil {
		log.Fatalf(`Unknow migration name "%s"`, cmd)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		return selected.Up(tx)
	})

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migration completed")
}
