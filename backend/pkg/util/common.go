package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
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
	return strconv.Itoa(num)
}

func GenerateTransactionNumber(orderNumber string) string {
	return fmt.Sprintf("%s_%s", uuid.NewString(), orderNumber)
}

func HashSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func JSONParse(data any) []byte {
	b, _ := json.Marshal(data)
	return b
}
