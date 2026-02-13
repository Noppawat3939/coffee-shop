package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	DataNotFound = "data not found"
	BodyInvalid  = "body invalid"
	Unauthorized = "unauthorized"
	Conflict     = "data conflict"
)

func Success(c *gin.Context, data ...interface{}) {
	res := gin.H{"code": http.StatusOK}

	if len(data) > 0 {
		res["data"] = data[0]
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
	res := gin.H{"code": http.StatusNotFound, "message": DataNotFound}

	c.JSON(http.StatusNotFound, res)
}

func ErrorBodyInvalid(c *gin.Context) {
	res := gin.H{"code": http.StatusBadRequest, "message": BodyInvalid}

	c.JSON(http.StatusBadRequest, res)
}

func ErrorUnauthorized(c *gin.Context) {
	res := gin.H{"code": http.StatusUnauthorized, "message": Unauthorized}

	c.JSON(http.StatusBadRequest, res)
}

func ErrorConflict(c *gin.Context) {
	res := gin.H{"code": http.StatusConflict, "message": Conflict}

	c.JSON(http.StatusConflict, res)
}
