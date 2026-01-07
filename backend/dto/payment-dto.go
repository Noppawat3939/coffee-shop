package dto

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type CreatePaymentTransactionLogRequest struct {
	OrderNumber string `json:"order_number" binding:"required"`
}
