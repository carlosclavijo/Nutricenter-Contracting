package repositories

import (
	"context"
	"database/sql"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrators"
	"github.com/google/uuid"
)

type AdministratorRepository struct {
	Db *sql.DB
}

func NewAdministratorRepository(db *sql.DB) administrators.AdministratorRepository {
	return &AdministratorRepository{Db: db}
}

func (r *AdministratorRepository) GetList(ctx context.Context) (*[]administrators.Administrator, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdministratorRepository) GetById(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdministratorRepository) Create(ctx context.Context, administrator *administrators.Administrator) error {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO administrators(id, name, phone) VALUES($1, $2, $3)", administrator.Id, administrator.Name, administrator.Phone)
	return err
}

func (r *AdministratorRepository) Update(ctx context.Context, administrator *administrators.Administrator) error {
	//TODO implement me
	panic("implement me")
}

func (r *AdministratorRepository) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
