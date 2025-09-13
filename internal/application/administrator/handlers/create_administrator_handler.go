package administratorhandlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"log"
)

func (h *AdministratorHandler) HandleCreate(ctx context.Context, cmd administratorcommands.CreateAdministratorCommand) (*administrators.Administrator, error) {
	adminFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password, cmd.Birth, cmd.Phone)

	if err != nil {
		log.
		return nil, err
	}

	admin, err := h.repository.Create(ctx, adminFactory)

	if err != nil {
		return nil, err
	}

	return admin, nil
}
