package valueobjects

import (
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

			assert.Equal(t, password.String(), tc.password)
			assert.True(t, isStrong)

			assert.Nil(t, err)
			assert.NoError(t, err)
		})
	}
}

func TestNewPassword_EmptyError(t *testing.T) {
	emp, err := NewPassword("")

	assert.NotNil(t, err)

	assert.ErrorIs(t, err, ErrEmptyPassword)

	assert.Empty(t, emp)
}

func TestNewPassword_LongError(t *testing.T) {
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

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrLongPassword)

			assert.Empty(t, password)
		})
	}
}

func TestNewPassword_ShortError(t *testing.T) {
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

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrShortPassword)

			assert.Empty(t, password)
		})
	}

}

func TestNewPassword_SoftError(t *testing.T) {
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

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrSoftPassword)

			assert.Empty(t, password)
		})
	}
}

func TestNewHashedPassword(t *testing.T) {
	cases := []struct {
		name, hash string
	}{
		{"Case 1", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa"},
		{"Case 2", "$2a$10$8YgR2mZc5WnQv1BfE0sHkQ7uLp9xC3tJ6rUoS4yVb2pN7zRaD0fGh"},
		{"Case 3", "$2a$10$F1sK8nVw3QzYp4LbH6rTqM2cXe9oJ5uSa7dG0vWb8nC1yZpR4tUqq"},
		{"Case 4", "$2a$10$wQ7kN2vL9sHf6R3bXtPzM1yGc4oU8eJa5rD0pZs2lVq6nYbF3uCwg"},
		{"Case 5", "$2a$10$zP4qT8nV1sLk3HcY6rFvM9wXe2oJ5uSa7dG0bWn8pC1yZrR4tUohw"},
		{"Case 6", "$2a$10$C7sL1pQ9vM2wHk8nXrFzG3yTq4oU6eJa5bD0pZs2lVq6nYbF3uCwe"},
		{"Case 7", "$2a$10$hR3vN6kL9sF2qPzM1yGdW8oXe4tU7eJa5bC0pZs2lVq6nYbF3uCwy"},
		{"Case 8", "$2a$10$qT9nV2sL4kH7rM1yFzG3wXe8oU6eJa5bD0pZs2lVq6nYbF3uCwPzW"},
		{"Case 9", "$2a$10$L2kH7nV3sQ9rM1yGdFzW8oXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQp"},
		{"Case 10", "$2a$10$M1yGdN7kL3sF2qPz9rH8wXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQps"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := NewHashedPassword(tc.hash)

			assert.NotEmpty(t, hash)
			assert.NoError(t, err)

			assert.Equal(t, hash.String(), tc.hash)

			assert.Nil(t, err)
		})
	}
}

func TestNewHashedPassword_EmptyError(t *testing.T) {
	emp, err := NewHashedPassword("")

	assert.NotNil(t, err)

	assert.ErrorIs(t, err, ErrEmptyHashedPassword)

	assert.Empty(t, emp)
}

func TestNewHashedPassword_FormatError(t *testing.T) {
	cases := []struct {
		name, hash string
	}{
		{"Case 1", "3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa"},
		{"Case 2", "8YgR2mZc5WnQv1BfE0sHkQ7uLp9xC3tJ6rUoS4yVb2pN7zRaD0fGh"},
		{"Case 3", "F1sK8nVw3QzYp4LbH6rTqM2cXe9oJ5uSa7dG0vWb8nC1yZpR4tUqq"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := NewHashedPassword(tc.hash)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrInvalidHashedPassword)

			assert.Empty(t, hash)
		})
	}
}

func TestNewHashedPassword_LengthError(t *testing.T) {
	cases := []struct {
		name, hash string
	}{
		{"Case 1", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1P"},
		{"Case 2", "$2a$10$8YgR2mZc5WnQv1BfE0sHkQ7uLp9xC3tJ6rUoS4yVb2pN7zRaD0fG"},
		{"Case 3", "$2a$10$F1sK8nVw3QzYp4LbH6rTqM2cXe9oJ5uSa7dG0vWb8nC1yZpR4tUq"},
		{"Case 4", "$2a$10$wQ7kN2vL9sHf6R3bXtPzM1yGc4oU8eJa5rD0pZs2lVq6nYbF3uCw"},
		{"Case 5", "$2a$10$zP4qT8nV1sLk3HcY6rFvM9wXe2oJ5uSa7dG0bWn8pC"},
		{"Case 6", "$2a$10$C7sL1pQ9vM2wHk8nXrFzG3yTq4oU6eJa5bD0pZs2lV"},
		{"Case 7", "$2a$10$hR3vN6kL9sF2qPzM1yGdW8oXe4tU7eJa5bC0pZs2lVq6nYbF3uCwywQ7kN2vL9sHf6R3bXtPzM1yGc4oU8eJa5rD0pZs2lVq6nYbF3uCw"},
		{"Case 8", "$2a$10$qT9nV2sL4kH7rM1yFzG3wXe8oU6eJa5bD0pZs2lVq6nYbF3uCwPzWwQ7kN2vL9sHf6R3bXtPzM1yGc4oU8eJa5rD0pZs2lVq6nYbF3uCw"},
		{"Case 9", "$2a$10$L2kH7nV3sQ9rM1yGdFzW8oXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQpWp3ñs"},
		{"Case 10", "$2a$10$M1yGdN7kL3sF2qPz9rH8wXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQpsfwegp"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := NewHashedPassword(tc.hash)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrLengthHashedPassword)

			assert.Empty(t, hash)
		})
	}
}
