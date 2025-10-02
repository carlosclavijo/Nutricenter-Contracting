package dto

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"time"
)

type PatientDTO struct {
	Id        string     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Gender    string     `json:"gender"`
	Birth     *time.Time `json:"birth,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
}

func MapToPatientDTO(patient *patients.Patient) *PatientDTO {
	if patient == nil {
		return nil
	}

	return &PatientDTO{
		Id:        patient.Id().String(),
		FirstName: patient.FirstName(),
		LastName:  patient.LastName(),
		Email:     patient.Email().Value(),
		Gender:    patient.Gender().String(),
		Birth:     patient.Birth().Value(),
		Phone:     patient.Phone().String(),
	}
}
