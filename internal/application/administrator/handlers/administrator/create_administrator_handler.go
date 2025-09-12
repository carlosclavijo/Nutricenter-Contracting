package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/service"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
)

type CreateAdministratorHandler struct {
	Service *service.AdministratorService
}

func (h *CreateAdministratorHandler) Handle(ctx context.Context, cmd command.CreateAdministratorCommand) (*administrators.Administrator, error) {
	return h.Service.Create(ctx, cmd.Name, cmd.Phone)
}
