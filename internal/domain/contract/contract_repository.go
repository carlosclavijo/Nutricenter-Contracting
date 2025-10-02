package contracts

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/google/uuid"
)

type ContractRepository interface {
	GetAll(ctx context.Context) ([]*Contract, error)
	GetById(ctx context.Context, id uuid.UUID) (*Contract, error)

	Create(ctx context.Context, contract *Contract) (*Contract, error)
	ChangeStatus(ctx context.Context, id uuid.UUID, status string) (*Contract, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)
	Count(ctx context.Context) (int, error)

	GetAllDeliveries(ctx context.Context) ([]*deliveries.Delivery, error)
	GetDeliveriesById(ctx context.Context, id uuid.UUID) (*deliveries.Delivery, error)

	UpdateDelivery(ctx context.Context, id uuid.UUID, delivery *deliveries.Delivery) (*deliveries.Delivery, error)
	ChangeStatusDelivery(ctx context.Context, id uuid.UUID, status string) (*deliveries.Delivery, error)
}
