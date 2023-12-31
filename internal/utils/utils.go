package utils

import (
	"regexp"
	"unicode"
)


func RegisterValidate(username string, email string, password string) bool {
	uV := usernameValidate(username)
	eV := emailValidate(email)
	pV := passwordValidate(password)

	if !uV || !eV || !pV {
		return false
	}
	return true
}

func usernameValidate(username string) bool {
	usenameRegex := regexp.MustCompile("^[a-zA-Z0-9]{5,30}$")
	return usenameRegex.MatchString(username)
}

func emailValidate(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func passwordValidate(pass string) bool {
	var (
		upp, low, num bool
		tot           uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || tot < 8 || tot > 30 {
		return false
	}

	return true
}
