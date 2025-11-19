package validation

import (
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"regexp"
	"time"
)

func ValidateProfileData(p models.Profile) error {
	now := time.Now()

	if !p.DateOfBirth.IsZero() {
		if p.DateOfBirth.After(now) {
			return fmt.Errorf("date of birth is in the future")
		}
		age := now.Year() - p.DateOfBirth.Year()
		if age > 130 {
			return fmt.Errorf("age cannot exceed 130 years")
		}
	}

	phoneRegex := regexp.MustCompile(`^[0-9+\-\(\) ]{7,20}$`)
	if p.Phone != "" && !phoneRegex.MatchString(p.Phone) {
		return fmt.Errorf("invalid phone format")
	}

	if len(p.Country) > 100 {
		return fmt.Errorf("invalid country")
	}

	if p.Age < 0 || p.Age > 120 {
		return fmt.Errorf("invalid age")
	}

	if len(p.Language) > 50 {
		return fmt.Errorf("invalid language")
	}

	return nil
}
