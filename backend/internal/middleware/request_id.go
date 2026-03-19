package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"
const RequestIDKeyHeader = "X-Request-ID"

func ReqID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(RequestIDKeyHeader)
		if reqID == "" {
			reqID = fmt.Sprintf("%s%s", "XRI-", uuid.NewString())
		}

		// set context for apply in middleware
		c.Set(RequestIDKey, reqID)

		// set to response headers
		c.Writer.Header().Set(RequestIDKeyHeader, reqID)

		c.Next()
	}
}
