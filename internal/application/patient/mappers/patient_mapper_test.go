package mappers

import (
	patients "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapToPatient_DTO_And_Response(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	email, _ := valueobjects.NewEmail("user@email.com")
	password, _ := valueobjects.NewPassword("Abc123!!")
	gender := valueobjects.Male
	birth, _ := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	phoneStr := "78978978"
	phone, _ := valueobjects.NewPhone(&phoneStr)

	patient := patients.NewPatient(firstName, lastName, email, password, gender, birth, phone)
	dto := MapToPatientDTO(patient)

	assert.NotNil(t, dto)

	assert.Equal(t, patient.Id().String(), dto.Id)
	assert.Equal(t, patient.FirstName(), dto.FirstName)
	assert.Equal(t, patient.LastName(), dto.LastName)
	assert.Equal(t, patient.Email().Value(), dto.Email)
	assert.Equal(t, patient.Gender().String(), dto.Gender)
	assert.Equal(t, patient.Birth().Value(), dto.Birth)
	assert.Equal(t, patient.Phone().String(), dto.Phone)

	response := MapToPatientResponse(dto, patient.LastLoginAt(), patient.CreatedAt(), patient.UpdatedAt(), patient.DeletedAt())

	assert.NotNil(t, response)

	assert.Exactly(t, *dto, response.PatientDTO)
	assert.Equal(t, patient.LastLoginAt(), response.LastLoginAt)
	assert.Equal(t, patient.CreatedAt(), response.CreatedAt)
	assert.Equal(t, patient.UpdatedAt(), response.UpdatedAt)
	assert.Equal(t, patient.DeletedAt(), response.DeletedAt)
}
