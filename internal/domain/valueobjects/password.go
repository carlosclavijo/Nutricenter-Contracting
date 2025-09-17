package valueobjects

import (
	"errors"
	"log"
	"regexp"
)

type Password struct {
	value string
}

func NewPassword(v string) (Password, error) {
	if v == "" {
		log.Printf("[valueobject:password] empty string")
		return Password{}, errors.New("password cannot be empty")
	} else if len(v) > 64 {
		log.Printf("[valueobject:password] password too long")
		return Password{}, errors.New("password is too long: maximum 64 characters")
	} else if len(v) < 8 {
		log.Printf("[valueobject:password] password too short")
		return Password{}, errors.New("password is too short: minimum 8 characters")
	} else if !isStrongPassword(v) {
		log.Printf("[valueobject:password] password too soft")
		return Password{}, errors.New("password isn't too strong")
	}
	return Password{v}, nil
}

func (p Password) String() string {
	return p.value
}

func isStrongPassword(v string) bool {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(v)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(v)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(v)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_\-+=<>?]`).MatchString(v)

	return hasLower && hasUpper && hasDigit && hasSpecial
}
