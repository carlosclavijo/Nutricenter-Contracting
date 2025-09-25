package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleUpdate(ctx context.Context, cmd commands.UpdateAdministratorCommand) (*administrators.Administrator, error) {
	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing email: %v", err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing password: %v", err)
		return nil, err
	}

	birth, err := valueobjects.NewBirthDate(cmd.Birth)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing birth: %v", err)
		return nil, err
	}

	phone, err := valueobjects.NewPhone(cmd.Phone)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing phone: %v", err)
		return nil, err
	}

	admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, cmd.Gender, birth, phone)
	if cmd.Id == uuid.Nil {
		log.Printf("[handler:administrator][HandleUpdate] Id '%v' is nil", cmd.Id)
		return nil, err
	}

	admin, err = h.repository.Update(ctx, admin)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error Updating Administrator: %v", err)
		return nil, err
	}

	return admin, nil
}
