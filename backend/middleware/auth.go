package middleware

import (
	"backend/util"
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
			util.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return

		}

		jwt := strings.TrimPrefix(authHeader, authPrefix)

		claims, err := util.ParseJWT(jwt)

		if err != nil {
			util.ErrorUnauthorized(c)
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
