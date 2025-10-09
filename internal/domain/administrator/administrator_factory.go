package administrators

import (
	"fmt"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"strings"
	"unicode"
)

type AdministratorFactory interface {
	Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*Administrator, error)
}

type administratorFactory struct{}

func (administratorFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*Administrator, error) {
	if firstName == "" {
		log.Printf("[factory:administrator] firstName '%s' is empty", firstName)
		return nil, ErrEmptyFirstNameAdministrator
	}

	if lastName == "" {
		log.Printf("[factory:administrator] lastName '%s' is empty", lastName)
		return nil, ErrEmptyLastNameAdministrator
	}

	if len(firstName) > 100 {
		log.Printf("[factory:administrator] firstName '%s' is too long, length %d, maximum is 100)", firstName, len(firstName))
		return nil, fmt.Errorf("%w: got %s, size %d", ErrLongFirstNameAdministrator, firstName, len(firstName))
	}

	if len(lastName) > 100 {
		log.Printf("[factory:administrator] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, fmt.Errorf("%w: got %s, size %d", ErrLongLastNameAdministrator, lastName, len(lastName))
	}

	if !isAlpha(firstName) {
		log.Printf("[factory:administrator] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, fmt.Errorf("%w: got %s", ErrNonAlphaFirstNameAdministrator, firstName)
	}

	if !isAlpha(lastName) {
		log.Printf("[factory:administrator] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, fmt.Errorf("%w: got %s", ErrNonAlphaLastNameAdministrator, lastName)
	}

	log.Printf("[factory:administrator][SUCCESS] administrator created")
	return NewAdministrator(firstName, lastName, email, password, gender, birth, phone), nil
}

func NewAdministratorFactory() AdministratorFactory {
	return &administratorFactory{}
}

func isAlpha(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	previousWasSpace := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			previousWasSpace = false
			continue
		} else if r == ' ' {
			if previousWasSpace {
				return false
			}
			previousWasSpace = true
		} else {
			return false
		}
	}

	return true
}
