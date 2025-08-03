package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)

	return value
}

func ParamToInt(c *gin.Context, key string) int {
	return ToInt(c.Param(key))

}
