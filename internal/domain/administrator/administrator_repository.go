package administrator

import (
	"context"
	"github.com/google/uuid"
)

type AdministratorRepository interface {
	GetList(ctx context.Context) (*[]Administrator, error)
	GetById(ctx context.Context, id uuid.UUID) (*Administrator, error)
	Create(ctx context.Context, administrator *Administrator) (*Administrator, error)
	Update(ctx context.Context, administrator *Administrator) (*Administrator, error)
	Delete(ctx context.Context, id uuid.UUID) (*Administrator, error)
}
