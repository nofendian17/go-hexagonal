package hash

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"user-svc/internal/shared/config"
)

const (
	saltSize = 16
)

type Hasher interface {
	GenerateRandomSalt() ([]byte, error)
	HashPassword(password string, salt []byte) string
	CheckPassword(hashedPassword, currentPassword string, salt []byte) bool
}

type hasher struct {
	config *config.Config
}

func NewHasher(config *config.Config) Hasher {
	return &hasher{
		config: config,
	}
}

func (h *hasher) GenerateRandomSalt() ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func (h *hasher) HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func (h *hasher) CheckPassword(hashedPassword, currentPassword string, salt []byte) bool {
	var currPasswordHash = h.HashPassword(currentPassword, salt)

	return hashedPassword == currPasswordHash
}
