package dto

import (
	"time"
)

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type CreateTxnLogRequest struct {
	OrderNumber string `json:"order_number" binding:"required"`
}

type CreateTxnResponse struct {
	TransactionNumber string    `json:"transaction_number"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	PaymentCode       string    `json:"payment_code"`
	ExpiredAt         time.Time `json:"expired_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type EnquireTxnRequst struct {
	TransactionNumber string `json:"transaction_number" binding:"required"`
	Status            string `json:"status"`
}

type EnquireTxnResponse struct {
	TransactionNumber string                      `json:"transaction_number"`
	Amount            float64                     `json:"amount"`
	Status            string                      `json:"status"`
	ExpiredAt         time.Time                   `json:"expired_at"`
	CreatedAt         time.Time                   `json:"created_at"`
	Order             EnquireTxnWithOrderResponse `json:"order"`
}

type EnquireTxnWithOrderResponse struct {
	ID          uint    `json:"id"`
	OrderNumber string  `json:"order_number"`
	Total       float64 `json:"total"`
	Status      string  `json:"status"`
}
