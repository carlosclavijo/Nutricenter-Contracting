package repository

import (
	"context"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
	"time"
)

func (r *Repository) Create(ctx context.Context, adm *administrators.Administrator) (*administrators.Administrator, error) {
	var (
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone, phoneVal                              *string
	)

	if adm.Phone() != nil {
		s := adm.Phone().String()
		phoneVal = s
	}

	query := `
		INSERT INTO administrator(id, first_name, last_name, email, password, gender, birth, phone)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(
		ctx, query, adm.Id, adm.FirstName, adm.LastName, adm.Email().Value(), adm.Password().String(), adm.Gender, adm.Birth, phoneVal,
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Create] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][Create] error concatenating administrator values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:administrator][Create] successfully created administrator in DB %v", admin)
	return admin, nil
}
