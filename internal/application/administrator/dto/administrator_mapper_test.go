package dto

import (
	administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapToAdministratorDTO(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	email, _ := valueobjects.NewEmail("user@email.com")
	password, _ := valueobjects.NewPassword("Abc123!!")
	gender := valueobjects.Male
	birth, _ := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	phoneStr := "78978978"
	phone, _ := valueobjects.NewPhone(&phoneStr)

	admin := administrators.NewAdministrator(firstName, lastName, email, password, gender, birth, phone)

	adminDb := MapToAdministratorDTO(*admin)

	assert.NotNil(t, adminDb)
	assert.Equal(t, adminDb.Id, admin.Id().String())
	assert.Equal(t, adminDb.FirstName, admin.FirstName())
	assert.Equal(t, adminDb.LastName, admin.LastName())
	assert.Equal(t, adminDb.Email, admin.Email().Value())
	assert.Equal(t, adminDb.Gender, admin.Gender().String())
	assert.Equal(t, adminDb.Birth, admin.Birth().Value())
	assert.Equal(t, adminDb.Phone, admin.Phone().String())
}
