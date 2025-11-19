package validation

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

func ValidateName(name string) error {
	if name == "" {
		return ErrInvalidName
	}

	if strings.Contains(name, " ") {
		return ErrInvalidName
	}

	if utf8.RuneCountInString(name) < 4 {
		return ErrInvalidName
	}

	if utf8.RuneCountInString(name) > 32 {
		return ErrInvalidName
	}

	return nil
}
