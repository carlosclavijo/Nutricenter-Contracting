package valueobjects

import (
	"fmt"
	"log"
	"unicode"
)

type Phone struct {
	value *string
}

func NewPhone(v *string) (*Phone, error) {
	if v == nil || *v == "" {
		return nil, nil
	} else if !isNumeric(*v) {
		log.Printf("[valueobject:phone] phone value isn't entirely numeric")
		return nil, fmt.Errorf("phone number '%s' isn't entirely numeric", *v)
	} else if len(*v) < 8 {
		log.Printf("[valueobject:phone] phone number too short")
		return nil, fmt.Errorf("phone number '%s' is too short('%d'), minimum length is 8", *v, len(*v))
	} else if len(*v) > 10 {
		log.Printf("[valueobject:phone] phone value too long")
		return nil, fmt.Errorf("phone number '%s' is too long('%d'), maximum length is 10", *v, len(*v))
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
