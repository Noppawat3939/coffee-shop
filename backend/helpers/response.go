package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data ...interface{}) {
	res := gin.H{"code": http.StatusOK}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, status int, msg string, data ...interface{}) {
	res := gin.H{"code": status, "message": msg}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(status, res)
}

func ErrorNotFound(c *gin.Context) {
	res := gin.H{"code": http.StatusNotFound, "message": "data not found"}

	c.JSON(http.StatusNotFound, res)
}
