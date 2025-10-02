package administrators

import (
	"context"
	"github.com/google/uuid"
)

type AdministratorRepository interface {
	GetAll(ctx context.Context) ([]*Administrator, error)
	GetList(ctx context.Context) ([]*Administrator, error)
	GetById(ctx context.Context, id uuid.UUID) (*Administrator, error)
	GetByEmail(ctx context.Context, email string) (*Administrator, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)

	Create(ctx context.Context, administrator *Administrator) (*Administrator, error)
	Update(ctx context.Context, administrator *Administrator) (*Administrator, error)
	Delete(ctx context.Context, id uuid.UUID) (*Administrator, error)
	Restore(ctx context.Context, id uuid.UUID) (*Administrator, error)

	CountAll(ctx context.Context) (int, error)
	CountActive(ctx context.Context) (int, error)
	CountDeleted(ctx context.Context) (int, error)
}
