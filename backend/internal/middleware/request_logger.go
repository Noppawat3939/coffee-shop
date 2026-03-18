package middleware

import (
	"backend/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ReqLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		reqID, _ := c.Get(RequestIDKey)

		status := c.Writer.Status()
		fields := []zap.Field{
			zap.Any("request_id", reqID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("duration", time.Since(time.Now())),
		}

		if status >= 500 {
			logger.Log.Error("[Error request]", fields...)
			return
		}

		logger.Log.Info("[Request]", fields...)
	}
}
