package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Print(err)
	return err == nil
}

func DummyCheck() {
	hash := "$2a$10$7EqJtq98hPqEX7fNZaFWoOeQF3bJ6X7kY3GZ0pQfM9vDOMkMt2Z4W"
	bcrypt.CompareHashAndPassword([]byte(hash), []byte("dummy"))
}
