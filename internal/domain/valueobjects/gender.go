package valueobjects

import "fmt"

type Gender string

const (
	Undefined Gender = "U"
	Male      Gender = "M"
	Female    Gender = "F"
)

func (g Gender) String() string {
	return string(g)
}

func ParseGender(s string) (Gender, error) {
	switch s {
	case "undefined":
		return Undefined, nil
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	default:
		return "", fmt.Errorf("invalid gender: %s", s)
	}
}
