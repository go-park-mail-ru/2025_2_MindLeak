package validation

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"

func ValidateUserData(u models.User) error {
	if err := ValidateName(u.Name); err != nil {
		return err
	}

	if u.Password != "" {
		if err := ValidatePassword(u.Password); err != nil {
			return err
		}
	}

	return nil
}
