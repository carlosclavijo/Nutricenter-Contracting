package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
)

func (h *PatientHandler) HandleUpdate(ctx context.Context, cmd commands.UpdatePatientCommand) (*patients.Patient, error) {
	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing email: %v", err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing password: %v", err)
		return nil, err
	}

	birth, err := valueobjects.NewBirthDate(cmd.Birth)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing birth: %v", err)
		return nil, err
	}

	phone, err := valueobjects.NewPhone(cmd.Phone)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing phone: %v", err)
		return nil, err
	}

	patient := patients.NewPatient(cmd.FirstName, cmd.LastName, email, password, cmd.Gender, birth, phone)

	if cmd.Id == uuid.Nil {
		log.Printf("[handler:patient][HandleUpdate] Id '%v' is nil", cmd.Id)
	}

	patient, err = h.repository.Update(ctx, patient)

	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error Updating Patient: %v", err)
		return nil, err
	}

	return patient, nil
}
