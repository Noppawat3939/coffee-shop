package helpers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)

	return value
}

func ParamToInt(c *gin.Context, key string) int {
	return ToInt(c.Param(key))

}

func IdStringToInts(s, splitter string) []int {
	var result []int

	for _, str := range strings.Split(s, splitter) {
		id := ToInt(str)

		result = append(result, id)
	}

	return result
}
