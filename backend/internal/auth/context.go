package auth

import (
	"backend/models"
	"backend/pkg/jwt"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) *models.UserJwyToken {
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

	return &models.UserJwyToken{
		ID:       uint(claims.EmployeeID),
		Username: claims.Username,
		Exp:      uint(claims.ExpiresAt.Time.Unix()),
	}
}
