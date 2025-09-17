package administratorhandlers

import "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"

type AdministratorHandler struct {
	repository administrators.AdministratorRepository
	factory    administrators.AdministratorFactory
}

func NewAdministratorHandler(r administrators.AdministratorRepository, f administrators.AdministratorFactory) *AdministratorHandler {
	return &AdministratorHandler{
		repository: r,
		factory:    f,
	}
}
