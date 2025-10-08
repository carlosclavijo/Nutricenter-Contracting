package patients

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPatientFromDB(t *testing.T) {
	phones := []string{"77141516", "77141517", "77141518", "77141519", "77141520"}
	deletedAts := []time.Time{time.Now(), time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, -2, 0), time.Now().AddDate(0, -3, 0), time.Now().AddDate(0, -4, 0)}
	cases := []struct {
		name, firstName, lastName, email, password, gender string
		birth, lastLoginAt, createdAt, updatedAt           time.Time
		phone                                              *string
		deletedAt                                          *time.Time
	}{
		{"Case 1", "Carlos", "Clavijo", "carlos@example.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "M", time.Now().AddDate(-25, 0, 0), time.Now(), time.Now(), time.Now(), &phones[0], nil},
		{"Case 2", "Ana", "Gomez", "ana@example.com", "$2a$10$8YgR2mZc5WnQv1BfE0sHkQ7uLp9xC3tJ6rUoS4yVb2pN7zRaD0fGh", "F", time.Now().AddDate(-30, 0, 0), time.Now(), time.Now(), time.Now(), nil, &deletedAts[0]},
		{"Case 3", "Luis", "Fernandez", "luis@example.com", "$2a$10$F1sK8nVw3QzYp4LbH6rTqM2cXe9oJ5uSa7dG0vWb8nC1yZpR4tUqq", "M", time.Now().AddDate(-22, 0, 0), time.Now(), time.Now(), time.Now(), &phones[1], &deletedAts[1]},
		{"Case 4", "Maria", "Lopez", "maria@example.com", "$2a$10$wQ7kN2vL9sHf6R3bXtPzM1yGc4oU8eJa5rD0pZs2lVq6nYbF3uCwg", "F", time.Now().AddDate(-28, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
		{"Case 5", "Juan", "Perez", "juan@example.com", "$2a$10$zP4qT8nV1sLk3HcY6rFvM9wXe2oJ5uSa7dG0bWn8pC1yZrR4tUohw", "M", time.Now().AddDate(-35, 0, 0), time.Now(), time.Now(), time.Now(), &phones[2], &deletedAts[2]},
		{"Case 6", "Lucia", "Martinez", "lucia@example.com", "$2a$10$C7sL1pQ9vM2wHk8nXrFzG3yTq4oU6eJa5bD0pZs2lVq6nYbF3uCwe", "F", time.Now().AddDate(-40, 0, 0), time.Now(), time.Now(), time.Now(), nil, &deletedAts[3]},
		{"Case 7", "Diego", "Ramirez", "diego@example.com", "$2a$10$hR3vN6kL9sF2qPzM1yGdW8oXe4tU7eJa5bC0pZs2lVq6nYbF3uCwy", "M", time.Now().AddDate(-23, 0, 0), time.Now(), time.Now(), time.Now(), &phones[3], nil},
		{"Case 8", "Sofia", "Torres", "sofia@example.com", "$2a$10$qT9nV2sL4kH7rM1yFzG3wXe8oU6eJa5bD0pZs2lVq6nYbF3uCwPzW", "F", time.Now().AddDate(-27, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
		{"Case 9", "Pedro", "Vargas", "pedro@example.com", "$2a$10$L2kH7nV3sQ9rM1yGdFzW8oXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQp", "M", time.Now().AddDate(-31, 0, 0), time.Now(), time.Now(), time.Now(), &phones[4], &deletedAts[4]},
		{"Case 10", "Elena", "Rojas", "elena@example.com", "$2a$10$M1yGdN7kL3sF2qPz9rH8wXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQps", "F", time.Now().AddDate(-26, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			patient, err := NewPatientFromDB(uuid.New(), tc.firstName, tc.lastName, tc.email, tc.password, tc.gender, tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt)

			assert.NoError(t, err)

			gender, err := valueobjects.ParseGender(tc.gender)

			assert.NotNil(t, patient)
			assert.NotNil(t, patient.Id())

			assert.Equal(t, tc.firstName, patient.FirstName())
			assert.Equal(t, tc.lastName, patient.LastName())
			assert.Equal(t, tc.email, patient.Email().Value())
			assert.Equal(t, tc.password, patient.Password().String())
			assert.Equal(t, gender.String(), patient.Gender().String())
			assert.Equal(t, tc.birth.Format(time.RFC3339), patient.Birth().Value().Format(time.RFC3339))
			assert.Equal(t, tc.lastLoginAt.Format(time.RFC3339), patient.LastLoginAt().Format(time.RFC3339))
			assert.Equal(t, tc.createdAt.Format(time.RFC3339), patient.CreatedAt().Format(time.RFC3339))
			assert.Equal(t, tc.updatedAt.Format(time.RFC3339), patient.UpdatedAt().Format(time.RFC3339))

			assert.NoError(t, err)

			if tc.phone != nil {
				assert.NotNil(t, patient.Phone)
				assert.Equal(t, *patient.Phone().String(), *tc.phone)
			} else {
				assert.Nil(t, patient.Phone())
			}

			if tc.deletedAt != nil {
				assert.NotNil(t, patient.DeletedAt())
				assert.Equal(t, tc.deletedAt.Format(time.RFC3339), patient.DeletedAt().Format(time.RFC3339))
			} else {
				assert.Nil(t, patient.DeletedAt())
			}

			tt := patient.LastLoginAt()
			time.Sleep(10 * time.Millisecond)
			patient.Logged()

			assert.True(t, patient.LastLoginAt().After(tt))
		})
	}
}

func TestNewAdministratorFromDB_Invalid(t *testing.T) {
	id := uuid.New()
	firstName := "Jane"
	lastName := "Does"
	email := "f@ilm@ail."
	password := "failedpassword"
	gender := "X"
	birth := time.Now()
	phone := "77887878A"
	lastLoginAt := time.Now().AddDate(0, 0, 10)
	createdAt := time.Now().AddDate(0, -6, 0)
	updatedAt := time.Now().AddDate(0, 3, 0)

	admin, err := NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, &phone, lastLoginAt, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrInvalidEmail)
	assert.Nil(t, admin)

	email = "john@doe.com"
	admin, err = NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, &phone, lastLoginAt, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrInvalidHashedPassword)
	assert.Nil(t, admin)

	password = "$2a$10$M1yGdN7kL3sF2qPz9rH8wXe4tU6eJa5bC0pZs2lVq6nYbF3uCwQps"
	admin, err = NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, &phone, lastLoginAt, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrNotAGender)
	assert.Nil(t, admin)

	gender = "male"
	admin, err = NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, &phone, lastLoginAt, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrUnderageDate)
	assert.Nil(t, admin)

	birth = time.Now().AddDate(-20, 0, 0)
	admin, err = NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, &phone, lastLoginAt, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrNotNumericPhoneNumber)
	assert.Nil(t, admin)
}
