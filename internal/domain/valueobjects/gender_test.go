package valueobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGender_String(t *testing.T) {
	u := Undefined
	m := Male
	f := Female
	other, _ := ParseGender("")

	assert.Equal(t, u.String(), "undefined")
	assert.Equal(t, m.String(), "male")
	assert.Equal(t, f.String(), "female")
	assert.Equal(t, other.String(), "unknown")
}
func TestParseGender(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected Gender
		wantErr  bool
	}{
		{"undefined word", "undefined", Undefined, false},
		{"undefined short", "U", Undefined, false},
		{"male word", "male", Male, false},
		{"male short", "M", Male, false},
		{"female word", "female", Female, false},
		{"female short", "F", Female, false},
		{"empty string", "", Gender(""), true},
		{"random invalid", "xyz", Gender(""), true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseGender(tc.input)

			if tc.wantErr {
				assert.NotNil(t, err)
				expected := fmt.Sprintf("invalid gender '%s'", Gender(tc.input))
				assert.ErrorContains(t, err, expected)
				assert.Equal(t, Gender(""), got)
			} else {
				assert.Nil(t, err)
				assert.NotEqual(t, Gender(""), got)
				assert.Exactly(t, tc.expected, got)
			}
		})
	}
}
