package dto

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type CreatePaymentTransactionLogRequest struct {
	OrderNumber string `json:"order_number" binding:"required"`
}

type EnquirPaymentTransactionLogRequst struct {
	TransactionNumber string `json:"transaction_number" binding:"required"`
	Status            string `json:"status"`
}
