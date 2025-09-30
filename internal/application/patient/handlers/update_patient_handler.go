package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
)

func (h *PatientHandler) HandleUpdate(ctx context.Context, cmd commands.UpdatePatientCommand) (*patients.Patient, error) {
	exist, err := h.repository.ExistById(ctx, cmd.Id)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error verifying if Patient exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:patient][HandleUpdate] the Patient doesn't exist '%v'", cmd.Id)
		return nil, errors.New("patient not found")
	}

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

	gender, err := valueobjects.ParseGender(cmd.Gender)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error parsing gender: %v", err)
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

	if cmd.Id == uuid.Nil {
		log.Printf("[handler:patient][HandleUpdate] Id '%v' is nil", cmd.Id)
		return nil, err
	}

	patient := patients.NewPatient(cmd.FirstName, cmd.LastName, email, password, gender, birth, phone)
	patient.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	patient, err = h.repository.Update(ctx, patient)
	if err != nil {
		log.Printf("[handler:patient][HandleUpdate] error Updating Patient: %v", err)
		return nil, err
	}

	return patient, nil
}
