package auth

import (
	"backend/internal/model"
	"backend/pkg/jwt"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) *model.UserJwyToken {
	user, exists := c.Get("user")

	if !exists {
		response.ErrorUnauthorized(c)
		c.Abort()
		return nil
	}

	claims, ok := user.(*jwt.JWTClaims)

	if !ok {
		response.ErrorUnauthorized(c)
		c.Abort()
		return nil
	}

	return &model.UserJwyToken{
		ID:       uint(claims.EmployeeID),
		Username: claims.Username,
		Role:     claims.Role,
		Exp:      uint(claims.ExpiresAt.Time.Unix()),
	}
}
