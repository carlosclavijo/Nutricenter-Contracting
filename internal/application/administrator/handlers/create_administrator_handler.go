package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *AdministratorHandler) HandleCreate(ctx context.Context, cmd commands.CreateAdministratorCommand) (*administrators.Administrator, error) {
	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error creating email '%s' object: %v", cmd.Email, err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error generating password: %v", err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(string(hashedPassword))
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error creating password object: %v", err)
		return nil, err
	}

	gender, err := valueobjects.ParseGender(cmd.Gender)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error creating gender object: %v", err)
		return nil, err
	}

	birth, err := valueobjects.NewBirthDate(cmd.Birth)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error creating birth date '%v' object: %v", birth, err)
		return nil, err
	}

	phone, err := valueobjects.NewPhone(cmd.Phone)
	if err != nil {
		log.Printf("[handler:administrator][HandleCreate] Error creating phone '%d' object: %v", phone, err)
		return nil, err
	}

	adminFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, email, password, gender, *birth, *phone)
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
