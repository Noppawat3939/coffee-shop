package dto

type QRRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}
