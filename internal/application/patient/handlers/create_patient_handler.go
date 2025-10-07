package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *PatientHandler) HandleCreate(ctx context.Context, cmd commands.CreatePatientCommand) (*dto.PatientResponse, error) {
	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error creating email '%s' object: %v", cmd.Email, err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error generating password: %v", err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(string(hashedPassword))
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error creating password object: %v", err)
		return nil, err
	}

	gender, err := valueobjects.ParseGender(cmd.Gender)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error creating gender object: %v", err)
		return nil, err
	}

	birth, err := valueobjects.NewBirthDate(cmd.Birth)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error creating birth date '%v' object: %v", birth, err)
		return nil, err
	}

	phone, err := valueobjects.NewPhone(cmd.Phone)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error creating phone '%d' object: %v", phone, err)
		return nil, err
	}

	patientFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, email, password, gender, birth, phone)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating AdministratorFactory: %v", err)
		return nil, err
	}

	patient, err := h.repository.Create(ctx, patientFactory)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating AdministratorRepository: %v", err)
		return nil, err
	}

	patientDto := mappers.MapToPatientDTO(patient)
	patientResponse := mappers.MapToPatientResponse(patientDto, patient.LastLoginAt(), patient.CreatedAt(), patient.UpdatedAt(), patient.DeletedAt())

	return patientResponse, nil
}
