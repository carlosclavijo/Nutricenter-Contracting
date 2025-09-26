package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *PatientHandler) HandleCreate(ctx context.Context, cmd commands.CreatePatientCommand) (*patients.Patient, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[handler:patient][HandleCreate] Error generating password: %v", err)
		return nil, err
	}

	patientFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, cmd.Email, string(password), cmd.Gender, cmd.Birth, cmd.Phone)

	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating PatientFactory: %v", err)
		return nil, err
	}

	admin, err := h.repository.Create(ctx, patientFactory)

	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating PatientRepository: %v", err)
		return nil, err
	}

	return admin, nil
}
