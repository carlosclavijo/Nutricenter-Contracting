package valueobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPassword(t *testing.T) {
	cases := []struct {
		name, password string
	}{
		{"Case 1", "Abcdef1!"},
		{"Case 2", "GoLangR0cks@"},
		{"Case 3", "P@ssw0rd123"},
		{"Case 4", "Str0ng#Key"},
		{"Case 5", "Test!2024A"},
		{"Case 6", "My_Secure9?"},
		{"Case 7", "XyZ@98765a"},
		{"Case 8", "SafePass!1"},
		{"Case 9", "ValidKey#2025"},
		{"Case 10", "Go!Dev123A"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			password, err := NewPassword(tc.password)
			isStrong := isStrongPassword(tc.password)

			assert.NotEmpty(t, password)
			assert.Nil(t, err)
			assert.NoError(t, err)

			assert.Equal(t, password.String(), tc.password)
			assert.True(t, isStrong)
		})
	}
}

func TestNewPassword_Invalid_Empty(t *testing.T) {
	emp, err := NewPassword("")

	assert.Empty(t, emp)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "password cannot be empty")
}

func TestNewPassword_Invalid_Long(t *testing.T) {
	cases := []struct {
		name, password string
	}{
		{"Case 1", "Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!"},
		{"Case 2", "GoLangR0cks@2025GoLangR0cks@2025GoLangR0cks@2025GoLangR0cks@2025s"},
		{"Case 3", "P@ssw0rd123P@ssw0rd123P@ssw0rd123P@ssw0rd123P@ssw0rd123P@ssw0rd123"},
		{"Case 4", "ValidKey#2025ValidKey#2025ValidKey#2025ValidKey#2025ValidKey#2025"},
		{"Case 5", "Test!2024ATest!2024ATest!2024ATest!2024ATest!2024ATest!2024ATest!2024A"},
		{"Case 6", "My_Secure9?My_Secure9?My_Secure9?My_Secure9?My_Secure9?My_Secure9?"},
		{"Case 7", "SafePass!1SafePass!1SafePass!1SafePass!1SafePass!1SafePass!1SafePass!1"},
		{"Case 8", "XyZ@98765aXyZ@98765aXyZ@98765aXyZ@98765aXyZ@98765aXyZ@98765aXyZ@98765a"},
		{"Case 9", "Go!Dev123AGo!Dev123AGo!Dev123AGo!Dev123AGo!Dev123AGo!Dev123AGo!Dev123A"},
		{"Case 10", "ComplexPass#99ComplexPass#99ComplexPass#99ComplexPass#99ComplexPass#99"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			password, err := NewPassword(tc.password)

			assert.Empty(t, password)
			assert.NotNil(t, err)

			expected := fmt.Sprintf("password '%s' is too long('%d') maximum 64 characters", tc.password, len(tc.password))
			assert.ErrorContains(t, err, expected)
		})
	}
}

func TestNewPassword_Invalid_Short(t *testing.T) {
	cases := []struct {
		name, password string
	}{
		{"Case 1", "Ab1!"},
		{"Case 2", "aA9#"},
		{"Case 3", "T3st!"},
		{"Case 4", "Qw1@"},
		{"Case 5", "P@5a"},
		{"Case 6", "Go!2"},
		{"Case 7", "aB3$"},
		{"Case 8", "1Aa!"},
		{"Case 9", "zZ9?"},
		{"Case 10", "tT8#"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			password, err := NewPassword(tc.password)

			assert.Empty(t, password)
			assert.NotNil(t, err)

			expected := fmt.Sprintf("password '%s' is too short('%d') minimum 8 characters", tc.password, len(tc.password))
			assert.ErrorContains(t, err, expected)
		})
	}

}

func TestNewPassword_Invalid_Soft(t *testing.T) {
	cases := []struct {
		name, password string
	}{
		{"Case 1", "abcdefg1"},
		{"Case 2", "ABCDEFG!"},
		{"Case 3", "12345678"},
		{"Case 4", "password!"},
		{"Case 5", "PASSWORD1"},
		{"Case 6", "NoSpecial1"},
		{"Case 7", "lowerUPPER"},
		{"Case 8", "@@@###$$$"},
		{"Case 9", "MixButNoNum"},
		{"Case 10", "Strongish1"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			password, err := NewPassword(tc.password)

			assert.Empty(t, password)
			assert.NotNil(t, err)
			expected := fmt.Sprintf("password '%s' isn't too strong", tc.password)
			assert.ErrorContains(t, err, expected)
		})
	}
}
