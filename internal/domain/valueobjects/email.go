package valueobjects

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

type Email struct {
	value string
}

const regex = `^[a-zA-Z0-9]([a-zA-Z0-9._%+\-]*[a-zA-Z0-9])?@([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}$`

var (
	ErrEmptyEmail   = errors.New("email cannot be empty")
	ErrInvalidEmail = errors.New("invalid email format")
	ErrLongEmail    = errors.New("email is too long, maximum size is 200")
)

func NewEmail(v string) (Email, error) {
	if v == "" {
		log.Printf("[valueobjects][email] empty string")
		return Email{}, ErrEmptyEmail
	}
	if !isValidEmail(v) {
		log.Printf("[valueobjects][email] Email '%s' is invalid", v)
		return Email{}, fmt.Errorf("%w: got %s", ErrInvalidEmail, v)
	} else if len(v) > 200 {
		log.Printf("[valueobjects][email] Email is too long")
		return Email{}, fmt.Errorf("%w: got %s, size %d", ErrLongEmail, v, len(v))
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
