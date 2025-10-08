package valueobjects

import (
	"errors"
	"fmt"
)

type Gender string

var ErrNotAGender error = errors.New("is not a gender")

const (
	Undefined Gender = "U" // Undefined
	Male      Gender = "M" // Male
	Female    Gender = "F" // Female
)

func (g Gender) String() string {
	switch g {
	case Undefined:
		return "undefined"
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

func ParseGender(s string) (Gender, error) {
	switch s {
	case "undefined", "U":
		return Undefined, nil
	case "male", "M":
		return Male, nil
	case "female", "F":
		return Female, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrNotAGender, s)
	}
}
