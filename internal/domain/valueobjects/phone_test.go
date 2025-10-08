package valueobjects

import (
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
			numeric := isNumeric(tc.phone)

			assert.NotEmpty(t, phone)
			assert.NoError(t, err)

			assert.Equal(t, *phone.String(), tc.phone)
			assert.True(t, numeric)

			assert.Nil(t, err)
		})
	}
}

func TestNewPhone_Empty(t *testing.T) {
	phone, err := NewPhone(nil)
	phoneStr := ""
	phone2, err2 := NewPhone(&phoneStr)

	assert.NoError(t, err)
	assert.NoError(t, err2)

	assert.Empty(t, phone)
	assert.Empty(t, phone2)
	assert.Nil(t, err)
	assert.Nil(t, err2)
}

func TestNewPhone_NotNumericError(t *testing.T) {
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
			numeric := isNumeric(tc.phone)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrNotNumericPhoneNumber)
			assert.False(t, numeric)

			assert.Nil(t, phone)

		})
	}
}

func TestNewPhone_ShortError(t *testing.T) {
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

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrShortPhoneNumber)

			assert.Nil(t, phone)
		})
	}
}

func TestNewPhone_LongError(t *testing.T) {
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

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrLongPhoneNumber)

			assert.Nil(t, phone)
		})
	}
}
