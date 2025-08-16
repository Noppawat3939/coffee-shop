package main

import (
	c "backend/config"
	"backend/db"
	"backend/middleware"
	"backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cfg c.Config

func init() {
	cfg = c.Load()
	db.Connect(cfg)
}

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

func main() {
	cfg = c.Load()
	database := db.Connect(cfg)

	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Use(gin.Logger(), gin.Recovery())

	r.Use(middleware.SetupCORS())

	routes.SetupRoutes(r, database)

	fmt.Print("✅ Starting server in port ", cfg.ServerPort)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":" + cfg.ServerPort)
}
