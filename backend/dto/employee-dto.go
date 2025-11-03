package dto

type CreateEmployeeRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateEmployeeRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Name     *string `json:"name,omitempty"`
	Role     *string `json:"role,omitempty"`
	Active   bool    `json:"active,omitempty"`
}
