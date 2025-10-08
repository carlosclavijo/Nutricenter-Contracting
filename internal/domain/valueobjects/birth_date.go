package valueobjects

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type BirthDate struct {
	value time.Time
}

var (
	ErrFutureDate   = errors.New("birthdate cannot be in the future")
	ErrUnderageDate = errors.New("cannot be an underage person")
)

func NewBirthDate(v time.Time) (BirthDate, error) {
	if v.After(time.Now()) {
		log.Printf("[valueobject:birth_date] birthdate '%v' cannot be in the future", v)
		return BirthDate{}, fmt.Errorf("%w: %s", ErrFutureDate, v)
	} else if !isAnAdult(v) {
		log.Printf("[valueobject:birth_date] cannot be an underage")
		return BirthDate{}, fmt.Errorf("%w: %v", ErrUnderageDate, v.Format(time.DateTime))
	}
	return BirthDate{value: v}, nil
}

func (b BirthDate) Value() time.Time {
	return b.value
}

func isAnAdult(v time.Time) bool {
	yearsAgo := time.Now().AddDate(-18, 0, 0)
	return !v.After(yearsAgo)
}
