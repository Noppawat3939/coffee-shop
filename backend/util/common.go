package util

import (
	"backend/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)

	return value
}

func ParamToInt(c *gin.Context, key string) int {
	return ToInt(c.Param(key))
}

func StringsToInts(strs []string) []int {
	var result []int

	for _, s := range strs {

		n := ToInt(s)

		result = append(result, n)
	}

	return result
}

func IntToString(num int) string {
	return strconv.Itoa(int(num))
}

func GetUserFromHeader(c *gin.Context) *models.UserJwyToken {
	user, exits := c.Get("user")

	if !exits {
		ErrorUnauthorized(c)
		c.Abort()
		return nil
	}

	// build claims
	claims, ok := user.(*JWTClaims)

	if !ok {
		ErrorUnauthorized(c)
		c.Abort()
		return nil
	}

	return &models.UserJwyToken{
		ID:       uint(claims.EmployeeID),
		Username: claims.Username,
		Exp:      uint(claims.ExpiresAt.Time.Unix()),
	}
}

func GenerateTransactionNumber(orderNumber string) string {
	return fmt.Sprintf("%s_%s", uuid.NewString(), orderNumber)
}
