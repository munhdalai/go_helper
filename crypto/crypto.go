package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

// AESEncrypt нь AES-GCM ашиглан plaintext-ийг encrypt хийнэ.
// Key нь 16, 24, эсвэл 32 байт байх ёстой (AES-128, AES-192, AES-256).
// Буцаах утга: nonce + ciphertext.
func AESEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return aesGCM.Seal(nonce, nonce, plaintext, nil), nil
}

// AESDecrypt нь AES-GCM ашиглан ciphertext-ийг decrypt хийнэ.
// Ciphertext нь nonce + encrypted data байх ёстой.
func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext хэтэрхий богино байна")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}

// deriveKey нь SHA-256 ашиглан string key-г 32 байт болгоно.
func deriveKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// AESEncryptString нь string plaintext-ийг encrypt хийж hex string буцаана.
// Key-г SHA-256-аар 32 байт болгоно (AES-256).
func AESEncryptString(plaintext, key string) (string, error) {
	encrypted, err := AESEncrypt([]byte(plaintext), deriveKey(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encrypted), nil
}

// AESDecryptString нь hex encoded ciphertext-ийг decrypt хийж string буцаана.
func AESDecryptString(ciphertext, key string) (string, error) {
	data, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	decrypted, err := AESDecrypt(data, deriveKey(key))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// HMACSign нь HMAC-SHA256 ашиглан мессежийг гарын үсэг зурна.
func HMACSign(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}

// HMACVerify нь HMAC-SHA256 гарын үсгийг баталгаажуулна (constant-time comparison).
func HMACVerify(message, key []byte, signature string) bool {
	expected := HMACSign(message, key)
	return hmac.Equal([]byte(expected), []byte(signature))
}
