package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/mappers"
	administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (h *AdministratorHandler) HandleLogin(ctx context.Context, cmd commands.LoginAdministratorCommand) (*dto.AdministratorResponse, error) {
	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing email '%s' %v", cmd.Email, err)
		return nil, err
	}

	exist, err := h.repository.ExistByEmail(ctx, cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleLogin] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleLogin] the Administrator doesn't exist '%v'", cmd.Email)
		return nil, administrators.ErrNotFoundAdministrator
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
		return nil, administrators.ErrInvalidCredentialsAdministrator
	}

	admin.Logged()
	admin, err = h.repository.Update(ctx, admin)
	if err != nil {
		log.Printf("[handler:administrator][HandleLogin] error Updating LastLoginAt of Administrator: %v", err)
		return nil, err
	}

	adminDto := mappers.MapToAdministratorDTO(admin)
	adminResponse := mappers.MapToAdministratorResponse(adminDto, admin.LastLoginAt(), admin.CreatedAt(), admin.UpdatedAt(), admin.DeletedAt())

	return adminResponse, nil
}
