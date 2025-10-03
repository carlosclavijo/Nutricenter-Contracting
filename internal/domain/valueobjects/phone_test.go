package valueobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPhone(t *testing.T) {
	cases := []struct {
		name, phone string
	}{
		{"Case 1", "71234567"},
		{"Case 2", "76543218"},
		{"Case 3", "70192834"},
		{"Case 4", "73456281"},
		{"Case 5", "74598126"},
		{"Case 6", "75643091"},
		{"Case 7", "72345618"},
		{"Case 8", "67891234"},
		{"Case 9", "70987653"},
		{"Case 10", "72139845"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			phone, err := NewPhone(&tc.phone)
			isNumeric := isNumeric(tc.phone)

			assert.NotEmpty(t, phone)
			assert.Nil(t, err)
			assert.NoError(t, err)

			assert.Equal(t, *phone.String(), tc.phone)
			assert.True(t, isNumeric)
		})
	}
}

func TestNewPhone_Nil(t *testing.T) {
	phone, err := NewPhone(nil)
	phoneStr := ""
	phone2, err2 := NewPhone(&phoneStr)

	assert.Nil(t, phone)
	assert.Nil(t, phone2)

	assert.NotNil(t, err)
	assert.NotNil(t, err2)

	assert.ErrorContains(t, err, "phone cannot be nil or empty")
	assert.ErrorContains(t, err2, "phone cannot be nil or empty")
}

func TestNewPhone_Invalid_IsNotNumeric(t *testing.T) {
	cases := []struct {
		name, phone string
	}{
		{"Case 1", "7654a3210"},
		{"Case 2", "60593b217"},
		{"Case 3", "78901x456"},
		{"Case 4", "723*91840"},
		{"Case 5", "61293q587"},
		{"Case 6", "75019z432"},
		{"Case 7", "690p38214"},
		{"Case 8", "70419#853"},
		{"Case 9", "78013n495"},
		{"Case 10", "69182w047"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			phone, err := NewPhone(&tc.phone)
			isNumeric := isNumeric(tc.phone)

			assert.Nil(t, phone)
			assert.NotNil(t, err)

			expected := fmt.Sprintf("phone number '%s' isn't entirely numeric", tc.phone)
			assert.ErrorContains(t, err, expected)
			assert.False(t, isNumeric)
		})
	}
}

func TestNewPhone_Invalid_Short(t *testing.T) {
	cases := []struct {
		name, phone string
	}{
		{"Case 1", "7123456"},
		{"Case 2", "7654321"},
		{"Case 3", "6987452"},
		{"Case 4", "7012345"},
		{"Case 5", "7345628"},
		{"Case 6", "7459812"},
		{"Case 7", "7564309"},
		{"Case 8", "7234561"},
		{"Case 9", "6789123"},
		{"Case 10", "7098765"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			phone, err := NewPhone(&tc.phone)

			assert.Nil(t, phone)
			assert.NotNil(t, err)

			expected := fmt.Sprintf("phone number '%s' is too short('%d'), minimum length is 8", tc.phone, len(tc.phone))
			assert.ErrorContains(t, err, expected)
		})
	}
}

func TestNewPhone_Invalid_Long(t *testing.T) {
	cases := []struct {
		name, phone string
	}{
		{"Case 1", "71234567890"},
		{"Case 2", "765432187654"},
		{"Case 3", "70123450987"},
		{"Case 4", "734562812345"},
		{"Case 5", "745981267890"},
		{"Case 6", "756430912345"},
		{"Case 7", "723456178901"},
		{"Case 8", "678912345678"},
		{"Case 9", "709876534210"},
		{"Case 10", "7123987456123"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			phone, err := NewPhone(&tc.phone)

			assert.Nil(t, phone)
			assert.NotNil(t, err)

			expected := fmt.Sprintf("phone number '%s' is too long('%d'), maximum length is 10", tc.phone, len(tc.phone))
			assert.ErrorContains(t, err, expected)
		})
	}
}
