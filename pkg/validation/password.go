package validation

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

func ValidatePassword(password string) error {
	if password == "" {
		return ErrInvalidPassword
	}

	if utf8.RuneCountInString(password) < 4 {
		return ErrInvalidPassword
	}

	if strings.Contains(password, " ") {
		return ErrInvalidPassword
	}

	if utf8.RuneCountInString(password) > 64 {
		return ErrInvalidPassword
	}

	return nil
}
