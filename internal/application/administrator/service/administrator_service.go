package service

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrators"
	"github.com/google/uuid"
)

type AdministratorService struct {
	repo    administrators.AdministratorRepository
	factory administrators.AdministratorFactory
}

func NewAdministratorService(repo administrators.AdministratorRepository, factory administrators.AdministratorFactory) *AdministratorService {
	return &AdministratorService{
		repo:    repo,
		factory: factory,
	}
}

func (a *AdministratorService) Create(ctx context.Context, name, phone string) (uuid.UUID, error) {
	admin, err := a.factory.Create(name, phone)
	if err != nil {
		return uuid.Nil, err
	} else if err = a.repo.Create(ctx, admin); err != nil {
		return uuid.Nil, err
	}
	return admin.Id, nil
}
