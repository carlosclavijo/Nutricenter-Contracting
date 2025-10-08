package valueobjects

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Password struct {
	value string
}

var (
	ErrEmptyPassword         = errors.New("password cannot be empty")
	ErrLongPassword          = errors.New("password is too long, maximum size is 64 characters")
	ErrShortPassword         = errors.New("password is too short, minimum size is 8 characters")
	ErrSoftPassword          = errors.New("password isn't strong enough")
	ErrEmptyHashedPassword   = errors.New("hashed password cannot be empty")
	ErrInvalidHashedPassword = errors.New("invalid hashed password format")
	ErrLengthHashedPassword  = errors.New("unexpected hashed password length")
)

func NewPassword(v string) (Password, error) {
	if v == "" {
		log.Printf("[valueobject:password] empty string")
		return Password{}, ErrEmptyPassword
	} else if len(v) > 64 {
		log.Printf("[valueobject:password] password too long")
		return Password{}, fmt.Errorf("%w: got %s, size %d", ErrLongPassword, v, len(v))
	} else if len(v) < 8 {
		log.Printf("[valueobject:password] password too short")
		return Password{}, fmt.Errorf("%w: got %s, size %d", ErrShortPassword, v, len(v))
	} else if !isStrongPassword(v) {
		log.Printf("[valueobject:password] password too soft")
		return Password{}, fmt.Errorf("%w: got %s", ErrSoftPassword, v)
	}
	return Password{v}, nil
}

func NewHashedPassword(v string) (Password, error) {
	if v == "" {
		log.Printf("[valueobject:password] empty hash")
		return Password{}, ErrEmptyHashedPassword
	}

	if !strings.HasPrefix(v, "$2") {
		log.Printf("[valueobject:password] invalid hash format")
		return Password{}, ErrInvalidHashedPassword
	}

	if len(v) != 60 {
		log.Printf("[valueobject:password] unexpected hash length: %d", len(v))
		return Password{}, fmt.Errorf("%w: size %d", ErrLengthHashedPassword, len(v))
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
