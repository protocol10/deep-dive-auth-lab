package validations

import (
	"errors"
	"unicode"
)

const (
	MinPasswordLength       = 8
	PasswordTooShortError   = "password must be at least 8 characters long"
	PasswordComplexityError = "password must contain at least one uppercase letter, one lowercase letter, one number, and one special character"
)

func IsPasswordLengthValid(password string) (bool, error) {
	if len(password) < MinPasswordLength {
		return false, errors.New(PasswordTooShortError)
	}
	return true, nil
}

func IsPasswordComplexityValid(password string) (bool, error) {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return false, errors.New(PasswordComplexityError)
	}

	return true, nil
}
