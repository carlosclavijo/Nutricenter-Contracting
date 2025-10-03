package valueobjects

import (
	"fmt"
	"log"
	"regexp"
)

const regex = `^[a-zA-Z0-9]([a-zA-Z0-9._%+\-]*[a-zA-Z0-9])?@([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}$`

type Email struct {
	value string
}

func NewEmail(v string) (Email, error) {
	if v == "" {
		log.Printf("[valueobjects][email] empty string")
		return Email{}, fmt.Errorf("email cannot be empty")
	}
	if !isValidEmail(v) {
		log.Printf("[valueobjects][email] Email '%s' is invalid", v)
		return Email{}, fmt.Errorf("email '%s' is an invalid email", v)
	} else if len(v) > 200 {
		log.Printf("[valueobjects][email] Email is too long")
		return Email{}, fmt.Errorf("email '%s' is too long ('%d'), max 200 characters", v, len(v))
	}
	return Email{value: v}, nil
}

func (e Email) Value() string {
	return e.value
}

func isValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(regex)
	return emailRegex.MatchString(email)
}
