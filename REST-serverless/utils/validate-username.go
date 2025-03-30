package utils

import (
	"errors"
	"unicode"
)

func ValidateUsername(username string) error {
	if len(username) < 6 {
		return errors.New("username must be at least 6 characters long")
	}

	for _, char := range username {
		if !unicode.IsLower(char) && !unicode.IsDigit(char) {
			return errors.New("username must be lowercase and alphanumeric only")
		}
	}

	return nil
}
