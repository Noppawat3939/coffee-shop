package middleware

import (
	"backend/pkg/jwt"
	"backend/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authPrefix = "Bearer "

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing token")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, authPrefix) {
			response.Error(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, authPrefix)

		claims, err := jwt.ParseJWT(token)
		if err != nil {
			response.ErrorUnauthorized(c)
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
