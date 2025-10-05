package dto

import (
	patients "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapToPatientDTO(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	email, _ := valueobjects.NewEmail("user@email.com")
	password, _ := valueobjects.NewPassword("Abc123!!")
	gender := valueobjects.Male
	birth, _ := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	phoneStr := "78978978"
	phone, _ := valueobjects.NewPhone(&phoneStr)

	patient := patients.NewPatient(firstName, lastName, email, password, gender, birth, phone)

	patientDb := MapToPatientDTO(*patient)

	assert.NotNil(t, patientDb)
	assert.Equal(t, patientDb.Id, patient.Id().String())
	assert.Equal(t, patientDb.FirstName, patient.FirstName())
	assert.Equal(t, patientDb.LastName, patient.LastName())
	assert.Equal(t, patientDb.Email, patient.Email().Value())
	assert.Equal(t, patientDb.Gender, patient.Gender().String())
	assert.Equal(t, patientDb.Birth, patient.Birth().Value())
	assert.Equal(t, patientDb.Phone, patient.Phone().String())
}
