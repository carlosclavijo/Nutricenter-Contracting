package valueobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBirthDate(t *testing.T) {
	past := time.Now().AddDate(-20, 0, 0)
	birthDate, err := NewBirthDate(past)

	assert.NotEmpty(t, birthDate)

	assert.Equal(t, past, birthDate.Value())

	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestNewBirthDate_Invalid_Future(t *testing.T) {
	future := time.Now().AddDate(5, 1, 2)
	birthDate, err := NewBirthDate(future)

	assert.NotNil(t, err)

	expected := fmt.Sprintf("birthdate '%v' cannot be in the future", future.Format("2006-01-02"))
	assert.ErrorContains(t, err, expected)

	assert.Empty(t, birthDate)
}

func TestNewBirthDate_Invalid_Underage(t *testing.T) {
	underage := time.Now().AddDate(-15, -10, -2)
	birthDate, err := NewBirthDate(underage)

	assert.NotNil(t, err)

	expected := fmt.Sprintf("date '%v' is an age of an underage person", underage.Format("2006-01-02"))
	assert.ErrorContains(t, err, expected)

	assert.Empty(t, birthDate)
}

func TestIsAnAdult(t *testing.T) {
	adult := time.Now().AddDate(-20, 0, 0)
	underage := time.Now().AddDate(-17, 9, 9)

	isAdult := isAnAdult(adult)
	isUnderage := isAnAdult(underage)

	assert.True(t, isAdult)
	assert.False(t, isUnderage)
}
