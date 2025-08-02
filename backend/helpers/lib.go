package helpers

import "strconv"

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)

	return value
}
