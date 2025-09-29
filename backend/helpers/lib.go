package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func toInt(s string) int {
	value, _ := strconv.Atoi(s)

	return value
}

func ParamToInt(c *gin.Context, key string) int {
	return toInt(c.Param(key))
}

func StringsToInts(strs []string) []int {
	var result []int

	for _, s := range strs {
		n := toInt(s)

		result = append(result, n)
	}

	return result
}

func IntToString(num int) string {
	return strconv.Itoa(int(num))
}
