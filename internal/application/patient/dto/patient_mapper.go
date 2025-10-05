package dto

import patients "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"

func MapToPatientDTO(patient patients.Patient) *PatientDTO {
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
