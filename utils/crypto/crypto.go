package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

// GenerateSalt generates a random salt
func GenerateSalt() string {
	saltBytes := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		log.Fatal(err)
	}
	salt := make([]byte, 32)
	hex.Encode(salt, saltBytes)
	return string(salt)
}

// HashPassword hashes a string
func HashPassword(password, salt string) string {
	hashedPasswordBytes, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		log.Fatal("Unable to hash password")
	}
	hashedPassword := make([]byte, 64)
	hex.Encode(hashedPassword, hashedPasswordBytes)
	return string(hashedPassword)
}

func GenerateToken() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	str := base64.URLEncoding.EncodeToString(b)
	return str, nil
}
