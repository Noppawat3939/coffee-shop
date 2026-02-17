package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	EmployeeID uint   `json:"employee_id"`
	Username   string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(employeeID uint, username string, exp time.Time) (string, error) {
	now := time.Now()

	signKey := (os.Getenv("JWT_SECRET"))
	claims := JWTClaims{
		EmployeeID: employeeID,
		Username:   username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "access_token",
		},
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tk.SignedString([]byte(signKey))
}

func ParseJWT(tkStr string) (*JWTClaims, error) {
	signKey := []byte(os.Getenv("JWT_SECRET"))

	tk, err := jwt.ParseWithClaims(tkStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}

	c, ok := tk.Claims.(*JWTClaims)
	if !ok || !tk.Valid {
		return nil, err
	}

	return c, nil
}
