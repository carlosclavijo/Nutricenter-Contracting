package deliveries

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/google/uuid"
)

type DeliveryRepository interface {
	GeList(ctx context.Context) (*[]dto.DeliveryDTO, error)
	GetById(ctx context.Context, id string) (*dto.DeliveryDTO, error)

	Create(ctx context.Context, contract *Delivery) (*Delivery, error)
	Update(ctx context.Context, contract *Delivery) (*Delivery, error)
	Delete(ctx context.Context, id uuid.UUID) (*Delivery, error)
	Restore(ctx context.Context, id uuid.UUID) (*Delivery, error)
}
