package administrators

import (
	"errors"
	valueobjects2 "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"time"
	"unicode"
)

type AdministratorFactory interface {
	Create(firstName, lastName, email, password, gender string, birth *time.Time, phone *string) (*Administrator, error)
}

type administratorFactory struct{}

func (a administratorFactory) Create(firstName, lastName, email, password, gender string, birth *time.Time, phone *string) (*Administrator, error) {
	if firstName == "" {
		log.Printf("[factory:administrator] firstName '%s' is empty", firstName)
		return nil, errors.New("firstName is empty")
	}
	if lastName == "" {
		log.Printf("[factory:administrator] lastName '%s' is empty", lastName)
		return nil, errors.New("lastName is empty")
	}
	if gender == "" {
		log.Printf("[factory:administrator] gender '%s' is empty", gender)
		return nil, errors.New("gender is empty")
	}

	emailVO, err := valueobjects2.NewEmail(email)
	if err != nil {
		log.Printf("[factory:administrator] Error creating email '%s' object: %v", email, err)
		return nil, err
	}
	passwordVO, err := valueobjects2.NewPassword(password)
	if err != nil {
		log.Printf("[factory:administrator] Error creating password object: %v", err)
		return nil, err
	}
	birthVO, err := valueobjects2.NewBirthDate(birth)
	if err != nil {
		log.Printf("[factory:administrator] Error creating birth date '%v' object: %v", birth, err)
		return nil, err
	}
	phoneVO, err := valueobjects2.NewPhone(phone)
	if err != nil {
		log.Printf("[factory:administrator] Error creating phone '%s' object: %v", phone, err)
		return nil, err
	}

	if len(firstName) > 100 {
		log.Printf("[factory:administrator] firstName '%s' is too long (length %d, maximum is 100)", firstName, len(firstName))
		return nil, errors.New("firstName is too long: maximum length is 100 characters")
	}
	if len(lastName) > 100 {
		log.Printf("[factory:administrator] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, errors.New("lastName is too long: maximum length is 100 characters")
	}
	if len(gender) != 1 {
		log.Printf("[factory:administrator] gender: '%v' can only be one character", gender)
		return nil, errors.New("gender isn't only one character")
	}
	if gender != "M" && gender != "F" && gender != "U" {
		log.Printf("[factory:administrator] gender: '%v' can only be 'M', 'F' or 'U'", gender)
		return nil, errors.New("gender is only 'M', 'F' or 'U'")
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
	return NewAdministrator(firstName, lastName, emailVO, passwordVO, gender, birthVO, phoneVO), nil
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
