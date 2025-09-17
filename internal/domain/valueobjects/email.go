package valueobjects

import (
	"errors"
	"log"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(v string) (Email, error) {
	if v == "" {
		log.Printf("[valueobjects][email] empty string")
		return Email{}, errors.New("email cannot be empty")
	}
	if !isValidEmail(v) {
		log.Printf("[valueobjects][email] Email is invalid")
		return Email{}, errors.New("invalid email")
	} else if len(v) > 200 {
		log.Printf("[valueobjects][email] Email is too long")
		return Email{}, errors.New("email too long: max 200 characters")
	}
	return Email{value: v}, nil
}

func (e Email) Value() string {
	return e.value
}

func isValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
