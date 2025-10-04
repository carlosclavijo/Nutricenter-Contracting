package valueobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGender_String(t *testing.T) {
	undefined := Undefined
	male := Male
	female := Female
	unknown, err := ParseGender("X")

	assert.Equal(t, undefined.String(), "undefined")
	assert.Equal(t, male.String(), "male")
	assert.Equal(t, female.String(), "female")
	assert.Equal(t, unknown.String(), "unknown")

	assert.Equal(t, Undefined, undefined)
	assert.Equal(t, Male, male)
	assert.Equal(t, Female, female)
	assert.NotEqual(t, Gender(""), unknown.String())

	expected := fmt.Sprintf("input '%s' is not a gender", "X")
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expected)

	g, err := ParseGender("undefined")
	assert.NoError(t, err)
	assert.Equal(t, Undefined, g)

	g, err = ParseGender("U")
	assert.NoError(t, err)
	assert.Equal(t, Undefined, g)

	g, err = ParseGender("male")
	assert.NoError(t, err)
	assert.Equal(t, Male, g)

	g, err = ParseGender("M")
	assert.NoError(t, err)
	assert.Equal(t, Male, g)

	g, err = ParseGender("female")
	assert.NoError(t, err)
	assert.Equal(t, Female, g)

	g, err = ParseGender("F")
	assert.NoError(t, err)
	assert.Equal(t, Female, g)

	g, err = ParseGender("invalid")
	assert.Error(t, err)
	assert.Equal(t, Gender(""), g)
}
