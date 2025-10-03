package valueobjects

import (
	"fmt"
	"log"
	"regexp"
)

type Password struct {
	value string
}

func NewPassword(v string) (Password, error) {
	if v == "" {
		log.Printf("[valueobject:password] empty string")
		return Password{}, fmt.Errorf("password cannot be empty")
	} else if len(v) > 64 {
		log.Printf("[valueobject:password] password too long")
		return Password{}, fmt.Errorf("password '%s' is too long('%d') maximum 64 characters", v, len(v))
	} else if len(v) < 8 {
		log.Printf("[valueobject:password] password too short")
		return Password{}, fmt.Errorf("password '%s' is too short('%d') minimum 8 characters", v, len(v))
	} else if !isStrongPassword(v) {
		log.Printf("[valueobject:password] password too soft")
		return Password{}, fmt.Errorf("password '%s' isn't too strong", v)
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
