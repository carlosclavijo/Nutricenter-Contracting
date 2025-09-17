package valueobjects

import (
	"errors"
	"log"
	"unicode"
)

type Phone struct {
	value *string
}

func NewPhone(v *string) (*Phone, error) {
	if v == nil {
		return nil, nil
	} else if !isNumeric(*v) {
		log.Printf("[valueobject:phone] phone value isn't entirely numeric")
		return &Phone{}, errors.New("invalid phone value")
	} else if len(*v) > 10 {
		log.Printf("[valueobject:phone] phone value too long")
		return &Phone{}, errors.New("phone number too long: maximum length is 10")
	}
	return &Phone{value: v}, nil
}

func (p Phone) String() *string {
	if p.value == nil {
		return nil
	}
	return p.value
}

func isNumeric(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
