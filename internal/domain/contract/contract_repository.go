package contracts

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/google/uuid"
)

type ContractRepository interface {
	GetAll(ctx context.Context) (*[]dto.ContractDTO, error)
	GetById(ctx context.Context, id uuid.UUID) (*dto.ContractDTO, error)

	Create(ctx context.Context, contract *Contract) (*Contract, error)
	ChangeStatus(ctx context.Context, id uuid.UUID, status string) (*Contract, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)
	Count(ctx context.Context) (int, error)

	GetAllDeliveries(ctx context.Context) (*[]dto.DeliveryDTO, error)
	GetDeliveriesById(ctx context.Context, id uuid.UUID) (*[]dto.DeliveryDTO, error)

	UpdateDelivery(ctx context.Context, id uuid.UUID, delivery *dto.DeliveryDTO) (*dto.DeliveryDTO, error)
	ChangeStatusDelivery(ctx context.Context, id uuid.UUID, status string) (*dto.DeliveryDTO, error)
}
