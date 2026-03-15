package main

import (
	c "backend/config"
	"backend/internal/database"
	"backend/internal/server"
	"log"
)

var cfg c.Config

func main() {
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

	log.Println("✅ Starting server in port ", cfg.ServerPort)

	if err := s.Start(cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
