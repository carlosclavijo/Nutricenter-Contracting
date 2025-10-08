package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmail_Empty(t *testing.T) {
	email, err := NewEmail("")

	assert.Empty(t, email)
	assert.Equal(t, email.Value(), "")
	assert.ErrorContains(t, err, "email cannot be empty")
}

func TestNewEmail(t *testing.T) {
	cases := []struct {
		name, email string
	}{
		{"Case 1", "user@example.com"},
		{"Case 2", "user.name@example.com"},
		{"Case 3", "user_name@example.co.uk"},
		{"Case 4", "user-name+tag@example.io"},
		{"Case 5", "user123@example-domain.com"},
		{"Case 6", "x@example.com"},
		{"Case 7", "very.common@example.com"},
		{"Case 8", "disposable.style.email.with+symbol@example.com"},
		{"Case 9", "other.email-with-dash@example.com"},
		{"Case 10", "admin@mailserver1.tv"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			email, err := NewEmail(tc.email)
			isValid := isValidEmail(tc.email)

			assert.NotEmpty(t, email)

			assert.Equal(t, email.Value(), tc.email)
			assert.True(t, isValid)

			assert.NoError(t, err)
		})
	}
}

func TestNewEmail_InvalidError(t *testing.T) {
	cases := []struct {
		name, email string
	}{
		{"Case 1", "plainaddress"},
		{"Case 2", "@example.com"},
		{"Case 3", "user@.com"},
		{"Case 4", "user@com"},
		{"Case 5", "user@-example.com"},
		{"Case 6", "user@example..com"},
		{"Case 7", "user@.example.com"},
		{"Case 8", "user@exam_ple.com"},
		{"Case 9", "user@example.com (Joe Smith)"},
		{"Case 10", "just\not\right@example.com"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			email, err := NewEmail(tc.email)
			isValid := isValidEmail(tc.email)

			assert.NotNil(t, err)
			assert.False(t, isValid)

			assert.ErrorIs(t, err, ErrInvalidEmail)

			assert.Empty(t, email)
		})
	}
}

func TestNewEmail_LongEmailError(t *testing.T) {
	cases := []struct {
		name, email string
	}{
		{"LongLocal", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"},
		{"LongSubdomain", "user@subdomainaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"},
		{"LongLocalAndSubdomain", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.user@subdomainaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			email, err := NewEmail(tc.email)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrLongEmail)

			assert.Empty(t, email)
		})
	}
}
