package main

import (
	"backend/config"
	"backend/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	pool := db.Connect(cfg)
	defer pool.Close()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	fmt.Print("✅ Starting server in port ", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
