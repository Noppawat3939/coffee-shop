package main

import (
	c "backend/config"
	"backend/db"
	"backend/middleware"
	"backend/routes"
	"encoding/base64"
	"fmt"

	pp "github.com/Frontware/promptpay"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var cfg c.Config

func init() {
	cfg = c.Load()
	db.Connect(cfg)
}

// Request body struct
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

	r.POST("/promptpay-qr", func(c *gin.Context) {
		var req QRRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "amount is required"})
			return
		}

		mockPhone := "0855873984"
		payment := pp.PromptPay{
			PromptPayID: mockPhone,
			Amount:      req.Amount,
			OneTime:     true,
		}

		// Generate PromptPay QR string
		qrString, err := payment.Gen()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate QR"})
			return
		}

		png, err := qrcode.Encode(qrString, qrcode.Medium, 256)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed generate QR image"})
		}

		// Return image as blob
		qr := base64.StdEncoding.EncodeToString(png)
		c.JSON(200, gin.H{"code": 200, "data": qr})
	})

	fmt.Print("✅ Starting server in port ", cfg.ServerPort)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":" + cfg.ServerPort)
}
