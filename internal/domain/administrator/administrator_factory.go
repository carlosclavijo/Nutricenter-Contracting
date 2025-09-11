package administrator

import (
	"errors"
	"unicode"
)

type AdministratorFactory interface {
	Create(administratorName, administratorPhone string) (*Administrator, error)
}

type administratorFactory struct{}

func (a administratorFactory) Create(administratorName, administratorPhone string) (*Administrator, error) {
	if administratorName == "" {
		return nil, errors.New("the administratorName is empty")
	} else if len(administratorName) > 100 {
		return nil, errors.New("the administratorName is too long")
	} else if administratorPhone == "" {
		return nil, errors.New("the administratorPhone is empty")
	} else if len(administratorPhone) > 8 {
		return nil, errors.New("the administratorPhone is too long")
	} else if !IsNumeric(administratorPhone) {
		return nil, errors.New("the administratorPhone is not a real number")
	}
	return NewAdministrator(administratorName, administratorPhone), nil
}

func NewAdministratorFactory() AdministratorFactory {
	return &administratorFactory{}
}

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
