package helpers

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateUUID нь шинэ UUID (v4) үүсгэнэ.
func GenerateUUID() string {
	return uuid.New().String()
}

// HashPassword нь bcrypt ашиглан password hash хийнэ.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword нь password hash-тай тохирч байгаа эсэхийг шалгана.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
