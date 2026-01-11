package middleware

import (
	"backend/models"
	"backend/util"
	"bytes"
	"fmt"
	"net/http"
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

func IdempotencyMiddleware(db *gorm.DB, ttlMinutes int) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(ID_KEY)
		if key == "" {
			util.Error(c, http.StatusBadRequest, fmt.Sprintf("%s %s", ID_KEY, "requred"))

			c.Abort()
			return
		}

		endpoint := c.FullPath()

		var record models.IdempotencyKey
		err := db.Where("key = ? AND endpoint = ?", key, endpoint).
			First(&record).Error

		if err == nil {
			util.Error(c, http.StatusConflict, "duplicate request (idempotency key already used)")
			c.Abort()
			return
		}

		// Not found and then capture response
		writer := &responseCapture{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		// Save response
		record = models.IdempotencyKey{
			ID:         uuid.New(),
			Key:        key,
			Endpoint:   endpoint,
			Response:   writer.body.Bytes(),
			StatusCode: writer.Status(),
			ExpiredAt:  time.Now().Add(time.Duration(ttlMinutes) * time.Minute),
		}
		db.Create(&record)
	}
}
