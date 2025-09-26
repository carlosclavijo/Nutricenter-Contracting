package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (h *AdministratorHandler) HandleLogin(ctx context.Context, cmd commands.LoginAdministratorCommand) (*administrators.Administrator, error) {
	exist, err := h.repository.ExistByEmail(ctx, cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleLogin] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleLogin] the Administrator doesn't exist '%v'", cmd.Email)
		return nil, errors.New("administrator not found")
	}

	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing email '%s' %v", cmd.Email, err)
		return nil, err
	}

	password, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing password: %v", err)
		return nil, err
	}

	admin, err := h.repository.GetByEmail(ctx, email.Value())
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password().String()), []byte(password.String())); err != nil {
		log.Printf("[handler:administrator][HandleLogin] invalid credentials for email=%s", cmd.Email)
		return nil, errors.New("invalid credentials")
	}

	admin.LastLoginAt = time.Now()
	admin, err = h.repository.Update(ctx, admin)
	if err != nil {
		log.Printf("[handler:administrator][HandleLogin] error Updating LastLoginAt of Administrator: %v", err)
		return nil, err
	}

	return admin, nil
}
