package valueobjects

import (
	"errors"
	"log"
	"time"
)

type BirthDate struct {
	value *time.Time
}

func NewBirthDate(v *time.Time) (*BirthDate, error) {
	if v == nil {
		return nil, nil
	} else if v.After(time.Now()) {
		log.Printf("[valueobject:birth_date] birthdate cannot be in the future")
		return nil, errors.New("birthdate cannot be in the future")
	} else if !isAnAdult(*v) {
		log.Printf("[valueobject:birth_date] cannot be an underage")
		return &BirthDate{}, errors.New("isn't an adult")
	}
	return &BirthDate{value: v}, nil
}

func (b BirthDate) Value() *time.Time {
	if b.value == nil {
		return nil
	}
	return b.value
}

func isAnAdult(v time.Time) bool {
	yearsAgo := time.Now().AddDate(-18, 0, 0)
	return !v.After(yearsAgo)
}
