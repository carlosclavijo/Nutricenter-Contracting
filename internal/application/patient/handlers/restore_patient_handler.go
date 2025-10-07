package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/mappers"
	"github.com/google/uuid"
	"log"
)

func (h *PatientHandler) HandleRestore(ctx context.Context, id uuid.UUID) (*dto.PatientResponse, error) {
	if id == uuid.Nil {
		log.Printf("[handler:patient][HandleRestore] Id '%v' is nil", id)
		return nil, errors.New("the id is not valid")
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		log.Printf("[handler:patient][HandleRestore] error verifying if Patient exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:patient][HandleRestore] the Patient doesn't exists '%v'", id)
		return nil, errors.New("patient not found")
	}

	patient, err := h.repository.Restore(ctx, id)
	if err != nil {
		log.Printf("[handler:patient][HandleRestore] error Deleting Patient: %v", err)
		return nil, err
	}

	patientDto := mappers.MapToPatientDTO(patient)
	patientResponse := mappers.MapToPatientResponse(patientDto, patient.LastLoginAt(), patient.CreatedAt(), patient.UpdatedAt(), patient.DeletedAt())

	return patientResponse, nil
}
