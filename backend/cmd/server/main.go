package main

import (
	c "backend/config"
	"backend/internal/database"
	"backend/internal/server"
	"backend/pkg/logger"
	"fmt"
	"log"
)

var cfg c.Config

func main() {
	logger.Init()
	defer func() {
		if err := logger.Log.Sync(); err != nil {
			fmt.Println("Error logger sync:", err)
		}
	}()

	cfg = c.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// run migration
	if err := database.Migration(db); err != nil {
		panic(err)
	}

	s := server.New(db)

	if err := s.Start(cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
