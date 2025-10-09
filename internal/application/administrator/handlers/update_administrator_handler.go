package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleUpdate(ctx context.Context, cmd commands.UpdateAdministratorCommand) (*dto.AdministratorResponse, error) {
	var err error

	if cmd.Id == uuid.Nil {
		log.Printf("[handler:administrator][HandleUpdate] Id '%v' is nil", cmd.Id)
		return nil, administrators.ErrEmptyIdAdministrator
	}

	email, err := valueobjects.NewEmail(cmd.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing email '%s': %v", cmd.Email, err)
		return nil, err
	}

	exist, err := h.repository.ExistById(ctx, cmd.Id)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleUpdate] the Administrator doesn't exist '%v'", cmd.Id)
		return nil, administrators.ErrNotFoundAdministrator
	}

	password, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing password: %v", err)
		return nil, err
	}

	gender, err := valueobjects.ParseGender(cmd.Gender)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing gender: %v", err)
		return nil, err
	}

	birth, err := valueobjects.NewBirthDate(cmd.Birth)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing birth '%v': %v", cmd.Birth, err)
		return nil, err
	}

	phone, err := valueobjects.NewPhone(cmd.Phone)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error parsing phone '%v': %v", cmd.Phone, err)
		return nil, err
	}

	admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, gender, birth, phone)
	admin.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	admin, err = h.repository.Update(ctx, admin)
	if err != nil {
		log.Printf("[handler:administrator][HandleUpdate] error Updating Administrator: %v", err)
		return nil, err
	}

	adminDto := mappers.MapToAdministratorDTO(admin)
	adminResponse := mappers.MapToAdministratorResponse(adminDto, admin.LastLoginAt(), admin.CreatedAt(), admin.UpdatedAt(), admin.DeletedAt())

	return adminResponse, nil
}
