package middleware

import (
	"backend/pkg/jwt"
	"backend/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var authPrefix = "Bearer "

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		var msg string = ""

		if authHeader == "" {
			msg = "missing token"
		}

		if authHeader != "" && !strings.HasPrefix(authHeader, authPrefix) {
			msg = "invalid token"
		}

		if msg != "" {
			response.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return

		}

		jwtStr := strings.TrimPrefix(authHeader, authPrefix)

		claims, err := jwt.ParseJWT(jwtStr)

		if err != nil {
			response.ErrorUnauthorized(c)
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
