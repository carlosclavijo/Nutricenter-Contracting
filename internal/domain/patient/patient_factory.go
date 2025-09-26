package patients

import (
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"time"
	"unicode"
)

type PatientFactory interface {
	Create(firstName, lastName, email, password, gender string, birth *time.Time, phone *string) (*Patient, error)
}

type patientFactory struct{}

func (a patientFactory) Create(firstName, lastName, email, password, gender string, birth *time.Time, phone *string) (*Patient, error) {
	if firstName == "" {
		log.Printf("[factory:patient] firstName '%s' is empty", firstName)
		return nil, errors.New("firstName is empty")
	}
	if lastName == "" {
		log.Printf("[factory:patient] lastName '%s' is empty", lastName)
		return nil, errors.New("lastName is empty")
	}
	if gender == "" {
		log.Printf("[factory:patient] gender '%s' is empty", gender)
		return nil, errors.New("gender is empty")
	}

	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		log.Printf("[factory:patient] Error creating email '%s' object: %v", email, err)
		return nil, err
	}
	passwordVO, err := valueobjects.NewPassword(password)
	if err != nil {
		log.Printf("[factory:patient] Error creating password object: %v", err)
		return nil, err
	}
	birthVO, err := valueobjects.NewBirthDate(birth)
	if err != nil {
		log.Printf("[factory:patient] Error creating birth date '%v' object: %v", birth, err)
		return nil, err
	}
	phoneVO, err := valueobjects.NewPhone(phone)
	if err != nil {
		log.Printf("[factory:patient] Error creating phone '%d' object: %v", phone, err)
		return nil, err
	}

	if len(firstName) > 100 {
		log.Printf("[factory:patient] firstName '%s' is too long (length %d, maximum is 100)", firstName, len(firstName))
		return nil, errors.New("firstName is too long: maximum length is 100 characters")
	}
	if len(lastName) > 100 {
		log.Printf("[factory:patient] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, errors.New("lastName is too long: maximum length is 100 characters")
	}
	if len(gender) != 1 {
		log.Printf("[factory:patient] gender: '%v' can only be one character", gender)
		return nil, errors.New("gender isn't only one character")
	}
	if gender != "M" && gender != "F" && gender != "U" {
		log.Printf("[factory:patient] gender: '%v' can only be 'M', 'F' or 'U'", gender)
		return nil, errors.New("gender is only 'M', 'F' or 'U'")
	}
	if !isAlpha(firstName) {
		log.Printf("[factory:patient] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, errors.New("firstName contains non-alphabetic characters")
	}
	if !isAlpha(lastName) {
		log.Printf("[factory:patient] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, errors.New("lastName contains non-alphabetic characters")
	}
	log.Printf("[factory:patient][SUCCESS] patient created")
	return NewPatient(firstName, lastName, emailVO, passwordVO, gender, birthVO, phoneVO), nil
}

func NewPatientFactory() PatientFactory {
	return &patientFactory{}
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || r == ' ' {
			return false
		}
	}
	return true
}
