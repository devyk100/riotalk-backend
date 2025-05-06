package utils

import (
	"errors"
	"unicode"
)

func ValidatePassword(input string) error {
	if len(input) == 0 {
		return errors.New("input cannot be empty")
	}

	if len(input) < 8 {
		return errors.New("input must be at least 8 characters long")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range input {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("input must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("input must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("input must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("input must contain at least one special character")
	}

	return nil
}
