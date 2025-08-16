package controllers

import (
	"backend/dto"
	hlp "backend/helpers"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentController struct{}

func NewPaymentController() *paymentController {
	return &paymentController{}
}

func (pc *paymentController) GeneratePromptPayQR(c *gin.Context) {
	var req dto.QRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		hlp.ErrorBodyInvalid(c)
		return
	}

	qr, err := services.GeneratePromptPayQR(req.Amount)
	if err != nil {
		hlp.Error(c, http.StatusInternalServerError, "failed generate QR promptpay")
		return
	}

	hlp.Success(c, gin.H{"qr": qr})
}
