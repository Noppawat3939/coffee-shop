package middleware

import (
	"backend/models"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const ID_KEY = "X-Idempotency-Key"

type responseCapture struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func IdempotencyMiddleware(db *gorm.DB, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(ID_KEY)
		if key == "" {
			c.Next()
			return
		}

		endpoint := c.FullPath()

		var record models.IdempotencyKey
		err := db.Where("key = ? AND endpoint = ?", key, endpoint).
			First(&record).Error

		// 🔁 Found → return cached response
		if err == nil {
			c.JSON(record.StatusCode, record.Response)
			c.Abort()
			return
		}

		// ⏳ Not found → capture response
		writer := &responseCapture{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		c.Next()

		// Save response
		record = models.IdempotencyKey{
			ID:         uuid.New(),
			Key:        key,
			Endpoint:   endpoint,
			Response:   writer.body.Bytes(),
			StatusCode: writer.Status(),
			ExpiredAt:  time.Now().Add(ttl),
		}
		db.Create(&record)
	}
}
