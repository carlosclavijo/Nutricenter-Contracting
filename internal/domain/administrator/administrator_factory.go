package administrators

import (
	"errors"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"unicode"
)

type AdministratorFactory interface {
	Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone vo.Phone) (*Administrator, error)
}

type administratorFactory struct{}

func (a administratorFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone vo.Phone) (*Administrator, error) {
	if firstName == "" {
		log.Printf("[factory:administrator] firstName '%s' is empty", firstName)
		return nil, errors.New("firstName is empty")
	}

	if lastName == "" {
		log.Printf("[factory:administrator] lastName '%s' is empty", lastName)
		return nil, errors.New("lastName is empty")
	}

	if len(firstName) > 100 {
		log.Printf("[factory:administrator] firstName '%s' is too long (length %d, maximum is 100)", firstName, len(firstName))
		return nil, errors.New("firstName is too long: maximum length is 100 characters")
	}

	if len(lastName) > 100 {
		log.Printf("[factory:administrator] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, errors.New("lastName is too long: maximum length is 100 characters")
	}

	if !isAlpha(firstName) {
		log.Printf("[factory:administrator] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, errors.New("firstName contains non-alphabetic characters")
	}

	if !isAlpha(lastName) {
		log.Printf("[factory:administrator] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, errors.New("lastName contains non-alphabetic characters")
	}

	log.Printf("[factory:administrator][SUCCESS] administrator created")
	return NewAdministrator(firstName, lastName, email, password, gender, &birth, &phone), nil
}

func NewAdministratorFactory() AdministratorFactory {
	return &administratorFactory{}
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || r == ' ' {
			return false
		}
	}
	return true
}
