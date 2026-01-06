package dto

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type CreatePaymentTransactionLogRequest struct {
	OrderID int `json:"order_id" binding:"required"`
}
