package patients

import (
	"fmt"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"strings"
	"unicode"
)

type PatientFactory interface {
	Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*Patient, error)
}

type patientFactory struct{}

func (patientFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*Patient, error) {
	if firstName == "" {
		log.Printf("[factory:patient] firstName '%s' is empty", firstName)
		return nil, fmt.Errorf("firstName is empty")
	}

	if lastName == "" {
		log.Printf("[factory:patient] lastName '%s' is empty", lastName)
		return nil, fmt.Errorf("lastName is empty")
	}

	if len(firstName) > 100 {
		log.Printf("[factory:patient] firstName '%s' is too long (length %d, maximum is 100)", firstName, len(firstName))
		return nil, fmt.Errorf("firstName '%s' is too long('%d'), maximum length is 100 characters", firstName, len(firstName))
	}

	if len(lastName) > 100 {
		log.Printf("[factory:patient] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, fmt.Errorf("lastName '%s' is too long('%d'), maximum length is 100 characters", lastName, len(lastName))
	}

	if !isAlpha(firstName) {
		log.Printf("[factory:patient] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, fmt.Errorf("firstName '%s' contains non-alphabetic characters", firstName)
	}

	if !isAlpha(lastName) {
		log.Printf("[factory:patient] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, fmt.Errorf("lastName '%s' contains non-alphabetic characters", lastName)
	}

	log.Printf("[factory:patient][SUCCESS] patient created")
	return NewPatient(firstName, lastName, email, password, gender, birth, phone), nil
}

func NewPatientFactory() PatientFactory {
	return &patientFactory{}
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
