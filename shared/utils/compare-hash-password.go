package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
