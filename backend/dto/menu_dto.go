package dto

type CreateMenuRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsAvailable bool   `json:"is_available"`
	Variations  []struct {
		Type  string  `json:"type" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Image string  `json:"image"`
	} `json:"variations"`
}
