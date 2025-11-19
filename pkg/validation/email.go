package validation

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

var emailRequired = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func ValidateEmail(email string) error {
	if email == "" {
		return ErrInvalidEmail
	}

	if !emailRequired.MatchString(email) {
		return ErrInvalidEmail
	}

	if utf8.RuneCountInString(email) > 320 {
		return ErrInvalidEmail
	}

	return nil
}
