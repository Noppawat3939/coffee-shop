package dto

type CreateOrderRequest struct {
	Customer   *string                 `json:"customer,omitempty"` // optional
	Variations []OrderVariationRequest `json:"variations" binding:"required,dive"`
	MemberID   uint                    `json:"member_id,omitempty"` //optional
}

type OrderVariationRequest struct {
	MenuVariationID uint `json:"menu_variation_id" binding:"required"`
	Amount          int  `json:"amount" binding:"required,min=1"`
}
