package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *PatientHandler) HandleLogin(ctx context.Context, cmd commands.LoginPatientCommand) (*dto.PatientResponse, error) {
	exist, err := h.repository.ExistByEmail(ctx, cmd.Email)
	if err != nil {
		log.Printf("[handler:patient][HandleLogin] error verifying if patient exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:patient][HandleLogin] the Patient doesn't exist '%v'", cmd.Email)
		return nil, errors.New("patient not found")
	}

	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing email '%s' %v", cmd.Email, err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing password: %v", err)
		return nil, err
	}

	patient, err := h.repository.GetByEmail(ctx, email.Value())
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(patient.Password().String()), []byte(password.String())); err != nil {
		log.Printf("[handler:patient][HandleLogin] invalid credentials for email=%s", cmd.Email)
		return nil, errors.New("invalid credentials")
	}

	patient.Logged()
	patient, err = h.repository.Update(ctx, patient)
	if err != nil {
		log.Printf("[handler:patient][HandleLogin] error Updating LastLoginAt of Patient: %v", err)
		return nil, err
	}

	patientDto := mappers.MapToPatientDTO(patient)
	patientResponse := mappers.MapToPatientResponse(patientDto, patient.LastLoginAt(), patient.CreatedAt(), patient.UpdatedAt(), patient.DeletedAt())

	return patientResponse, nil
}
