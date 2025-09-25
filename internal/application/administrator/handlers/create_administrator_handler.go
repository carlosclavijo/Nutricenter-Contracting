package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"log"
)

func (h *AdministratorHandler) HandleCreate(ctx context.Context, cmd commands.CreateAdministratorCommand) (*administrators.Administrator, error) {
	adminFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth, cmd.Phone)

	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] error Creating AdministratorFactory: %v", err)
		return nil, err
	}

	admin, err := h.repository.Create(ctx, adminFactory)

	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] error Creating AdministratorRepository: %v", err)
		return nil, err
	}

	return admin, nil
}
