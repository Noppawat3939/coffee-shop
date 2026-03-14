package middleware

import (
	"backend/internal/model"
	"backend/pkg/response"
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const IDKey = "X-Idempotency-Key"

type responseCapture struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseCapture) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func IdempotencyMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		key := c.GetHeader(IDKey)

		if key == "" {
			response.Error(c, http.StatusBadRequest, fmt.Sprintf("%s required", IDKey))
			c.Abort()
			return
		}

		endpoint := c.FullPath()

		var record model.IdempotencyKey

		err := db.Where("key = ? AND endpoint = ?", key, endpoint).
			First(&record).Error

		if err == nil {
			response.Error(c, http.StatusConflict, "duplicate request")
			c.Abort()
			return
		}

		writer := &responseCapture{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = writer

		c.Next()

		record = model.IdempotencyKey{
			ID:         uuid.New(),
			Key:        key,
			Endpoint:   endpoint,
			Response:   writer.body.Bytes(),
			StatusCode: writer.Status(),
			ExpiredAt:  time.Now().Add(2 * time.Minute),
		}

		db.Create(&record)
	}
}
