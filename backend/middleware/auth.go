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
		authHeader := util.GetAuthHeader(c)

		if authHeader == "" || !strings.HasPrefix(authHeader, authPrefix) {
			util.Error(c, http.StatusUnauthorized, "invalid token")
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
