package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGender(t *testing.T) {
	undefined := Undefined
	male := Male
	female := Female
	unknown, err := ParseGender("X")

	assert.NotNil(t, err)

	assert.Equal(t, undefined.String(), "undefined")
	assert.Equal(t, male.String(), "male")
	assert.Equal(t, female.String(), "female")
	assert.Equal(t, unknown.String(), "unknown")
	assert.Equal(t, Undefined, undefined)
	assert.Equal(t, Male, male)
	assert.Equal(t, Female, female)
	assert.NotEqual(t, Gender(""), unknown.String())

	assert.ErrorIs(t, err, ErrNotAGender)

	g, err := ParseGender("undefined")
	assert.Equal(t, Undefined, g)
	assert.NoError(t, err)

	g, err = ParseGender("U")
	assert.Equal(t, Undefined, g)
	assert.NoError(t, err)

	g, err = ParseGender("male")
	assert.Equal(t, Male, g)
	assert.NoError(t, err)

	g, err = ParseGender("M")
	assert.Equal(t, Male, g)
	assert.NoError(t, err)

	g, err = ParseGender("female")
	assert.Equal(t, Female, g)
	assert.NoError(t, err)

	g, err = ParseGender("F")
	assert.Equal(t, Female, g)
	assert.NoError(t, err)

	g, err = ParseGender("invalid")
	assert.Equal(t, Gender(""), g)
	assert.ErrorIs(t, err, ErrNotAGender)
}
