package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user-svc/internal/shared/config"
)

const (
	saltLength = 16
	bcryptCost = bcrypt.DefaultCost
)

var (
	secretKey = config.New().App.Key
)

// HashPassword hashes a given password with bcrypt algorithm and returns the hashed password and salt used
func HashPassword(password string) (string, string, error) {
	saltBytes := make([]byte, saltLength)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", err
	}
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password+salt+secretKey), bcryptCost)
	if err != nil {
		return "", "", err
	}
	hashedPassword := base64.StdEncoding.EncodeToString(hashedPasswordBytes)

	return hashedPassword, salt, nil
}

// CheckPassword checks if a given password matches the hashed password with the provided salt
func CheckPassword(password, hashedPassword, salt string) (bool, error) {
	hashedPasswordBytes, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, err
	}
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false, err
	}

	hash := sha256.New()
	hash.Write([]byte(password + string(saltBytes) + string(secretKey)))
	hashedPasswordBytes2, err := bcrypt.GenerateFromPassword(hash.Sum(nil), bcryptCost)
	if err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, hashedPasswordBytes2); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
