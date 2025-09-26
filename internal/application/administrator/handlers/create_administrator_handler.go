package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *AdministratorHandler) HandleCreate(ctx context.Context, cmd commands.CreateAdministratorCommand) (*administrators.Administrator, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error generating password: %v", err)
		return nil, err
	}

	adminFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, cmd.Email, string(password), cmd.Gender, cmd.Birth, cmd.Phone)

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
