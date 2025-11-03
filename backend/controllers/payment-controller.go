package controllers

import (
	"backend/dto"
	"backend/services"
	"backend/util"
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
		util.ErrorBodyInvalid(c)
		return
	}

	qr, err := services.GeneratePromptPayQR(req.Amount)
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "failed generate QR promptpay")
		return
	}

	util.Success(c, gin.H{"qr": qr})
}
