package security

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const saltLen = 16

const (
	time    = 1         // итерации
	memory  = 64 * 1024 // 64 МБ
	threads = 4
	keyLen  = 32
)

func HashPassword(plainPassword string) ([]byte, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	hashedPass := argon2.IDKey([]byte(plainPassword), salt, time, memory, threads, keyLen)

	return append(salt, hashedPass...), nil
}

func CheckPassword(storedHash []byte, plainPassword string) bool {
	if len(storedHash) < saltLen {
		return false
	}

	salt := storedHash[:saltLen]
	userHash := argon2.IDKey([]byte(plainPassword), salt, time, memory, threads, keyLen)

	expected := storedHash[saltLen:]
	return bytes.Equal(userHash, expected)
}
