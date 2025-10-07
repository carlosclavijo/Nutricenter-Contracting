package mappers

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
	gender := valueobjects.Male

	email, err := valueobjects.NewEmail("user@email.com")
	assert.NoError(t, err)

	password, err := valueobjects.NewPassword("Abc123!!")
	assert.NoError(t, err)

	birth, err := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	assert.NoError(t, err)

	phoneStr := "78978978"
	phone, err := valueobjects.NewPhone(&phoneStr)
	assert.NoError(t, err)

	admin := administrators.NewAdministrator(firstName, lastName, email, password, gender, birth, phone)
	dto := MapToAdministratorDTO(admin)

	assert.NotNil(t, dto)

	assert.Equal(t, admin.Id().String(), dto.Id)
	assert.Equal(t, admin.FirstName(), dto.FirstName)
	assert.Equal(t, admin.LastName(), dto.LastName)
	assert.Equal(t, admin.Email().Value(), dto.Email)
	assert.Equal(t, admin.Gender().String(), dto.Gender)
	assert.Equal(t, admin.Birth().Value(), dto.Birth)
	assert.Equal(t, admin.Phone().String(), dto.Phone)

	response := MapToAdministratorResponse(dto, admin.LastLoginAt(), admin.CreatedAt(), admin.UpdatedAt(), admin.DeletedAt())

	assert.NotNil(t, response)

	assert.Exactly(t, *dto, response.AdministratorDTO)
	assert.Equal(t, admin.LastLoginAt(), response.LastLoginAt)
	assert.Equal(t, admin.CreatedAt(), response.CreatedAt)
	assert.Equal(t, admin.UpdatedAt(), response.UpdatedAt)
	assert.Equal(t, admin.DeletedAt(), response.DeletedAt)

}
