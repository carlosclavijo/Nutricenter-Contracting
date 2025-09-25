package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"log"
)

func (h *PatientHandler) HandleCreate(ctx context.Context, cmd commands.CreatePatientCommand) (*patients.Patient, error) {
	adminFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth, cmd.Phone)

	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating PatientFactory: %v", err)
		return nil, err
	}

	admin, err := h.repository.Create(ctx, adminFactory)

	if err != nil {
		log.Printf("[handler:patient][HandleCreate] error Creating PatientRepository: %v", err)
		return nil, err
	}

	return admin, nil
}
