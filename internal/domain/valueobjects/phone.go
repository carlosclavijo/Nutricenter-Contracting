package valueobjects

import (
	"errors"
	"fmt"
	"log"
	"unicode"
)

type Phone struct {
	value *string
}

var (
	ErrNotNumericPhoneNumber = errors.New("phone number isn't entirely numeric")
	ErrShortPhoneNumber      = errors.New("phone number is too short, minimum size is 8")
	ErrLongPhoneNumber       = errors.New("phone nunmber is too long, maximum size is 10")
)

func NewPhone(v *string) (*Phone, error) {
	if v == nil || *v == "" {
		return nil, nil
	} else if !isNumeric(*v) {
		log.Printf("[valueobject:phone] phone value isn't entirely numeric")
		return nil, fmt.Errorf("%w: got %s", ErrNotNumericPhoneNumber, *v)
	} else if len(*v) < 8 {
		log.Printf("[valueobject:phone] phone number too short")
		return nil, fmt.Errorf("%w: got %s, size %d", ErrShortPhoneNumber, *v, len(*v))
	} else if len(*v) > 10 {
		log.Printf("[valueobject:phone] phone value too long")
		return nil, fmt.Errorf("%w, got %s, size %d", ErrLongPhoneNumber, *v, len(*v))
	}
	return &Phone{value: v}, nil
}

func (p Phone) String() *string {
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
