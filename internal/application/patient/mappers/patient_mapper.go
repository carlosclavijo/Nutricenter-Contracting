package mappers

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"time"
)

func MapToPatientDTO(patient *patients.Patient) *dto.PatientDTO {
	var phone *string
	if patient.Phone() != nil {
		p := patient.Phone().String()
		phone = p
	}

	return &dto.PatientDTO{
		Id:        patient.Id().String(),
		FirstName: patient.FirstName(),
		LastName:  patient.LastName(),
		Email:     patient.Email().Value(),
		Gender:    patient.Gender().String(),
		Birth:     patient.Birth().Value(),
		Phone:     phone,
	}
}

func MapToPatientResponse(patient *dto.PatientDTO, last, created, updated time.Time, deleted *time.Time) *dto.PatientResponse {
	return &dto.PatientResponse{
		PatientDTO:  *patient,
		LastLoginAt: last,
		CreatedAt:   created,
		UpdatedAt:   updated,
		DeletedAt:   deleted,
	}
}
