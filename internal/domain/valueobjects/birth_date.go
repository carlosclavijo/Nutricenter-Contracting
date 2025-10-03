package valueobjects

import (
	"fmt"
	"log"
	"time"
)

type BirthDate struct {
	value time.Time
}

func NewBirthDate(v time.Time) (BirthDate, error) {
	if v.After(time.Now()) {
		log.Printf("[valueobject:birth_date] birthdate '%v' cannot be in the future", v)
		return BirthDate{}, fmt.Errorf("birthdate '%v' cannot be in the future", v.Format("2006-01-02"))
	} else if !isAnAdult(v) {
		log.Printf("[valueobject:birth_date] cannot be an underage")
		return BirthDate{}, fmt.Errorf("date '%v' is an age of an underage person", v.Format("2006-01-02"))
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
