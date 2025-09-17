package administratorhandlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleUpdate(ctx context.Context, cmd administratorcommands.UpdateAdministratorCommand) (*administrators.Administrator, error) {
	email, _ := valueobjects.NewEmail(cmd.Email)
	password, _ := valueobjects.NewPassword(cmd.Password)
	birth, _ := valueobjects.NewBirthDate(cmd.Birth)
	phone, _ := valueobjects.NewPhone(cmd.Phone)

	admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, cmd.Gender, birth, phone)

	if cmd.Id == uuid.Nil {
		log.Printf("[handler:administrator][HandleUpdate] Id '%v' is nil", cmd.Id)
	}

	admin, err := h.repository.Update(ctx, admin)

	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error Updating Administrator: %v", err)
		return nil, err
	}

	return admin, nil
}
