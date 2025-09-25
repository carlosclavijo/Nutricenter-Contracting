package patients

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/google/uuid"
)

type PatientRepository interface {
	GetAll(ctx context.Context) (*[]dto.PatientDTO, error)
	GetList(ctx context.Context) (*[]dto.PatientDTO, error)
	GetById(ctx context.Context, id uuid.UUID) (*dto.PatientDTO, error)
	GetByEmail(ctx context.Context, email string) (*Patient, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)

	Create(ctx context.Context, patient *Patient) (*Patient, error)
	Update(ctx context.Context, patient *Patient) (*Patient, error)
	Delete(ctx context.Context, id uuid.UUID) (*Patient, error)
	Restore(ctx context.Context, id uuid.UUID) (*Patient, error)

	CountAll(ctx context.Context) (int, error)
	CountActive(ctx context.Context) (int, error)
	CountDeleted(ctx context.Context) (int, error)
}
