package utils

import "unicode"

func IsStrongPassword(p string) bool {
	if len(p) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit bool

	for _, c := range p {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
