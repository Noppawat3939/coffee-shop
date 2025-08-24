package dto

type GetMemberRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type CreateMemberRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	FullName    string `json:"full_name"`
}
