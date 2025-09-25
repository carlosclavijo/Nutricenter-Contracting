package contracts

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/google/uuid"
)

type ContractRepository interface {
	GetList(ctx context.Context) (*[]dto.ContractDTO, error)
	GetById(ctx context.Context, id uuid.UUID) (*dto.ContractDTO, error)

	Create(ctx context.Context, contract *Contract) (*Contract, error)
	Update(ctx context.Context, contract *Contract) (*Contract, error)
	Delete(ctx context.Context, id uuid.UUID) (*Contract, error)
	Restore(ctx context.Context, id uuid.UUID) (*Contract, error)
}
