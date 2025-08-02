package main

import (
	c "backend/config"
	"backend/db"
	"backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cfg c.Config

func init() {
	cfg = c.Load()
	db.Connect(cfg)
}

func main() {
	cfg = c.Load()
	database := db.Connect(cfg)
	r := gin.Default()

	api := r.Group("/api")

	routes.IntialMenuRoutes(api, database)

	fmt.Print("✅ Starting server in port ", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
