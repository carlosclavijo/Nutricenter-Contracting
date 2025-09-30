package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
	"time"
)

type AdministratorRepository struct {
	Db *sql.DB
}

func (r *AdministratorRepository) GetAll(ctx context.Context) (*[]dto.AdministratorDTO, error) {
	var admins []dto.AdministratorDTO

	query := `
		SELECT id, first_name, last_name, email, gender, birth, phone
		FROM administrator
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:administrator][GetAll] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("[repository:administrator][GetAll] failed to close rows: %v", err)
			return
		}
	}(rows)
	for rows.Next() {
		var admin dto.AdministratorDTO
		err = rows.Scan(&admin.Id, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Gender, &admin.Birth, &admin.Phone)
		if err != nil {
			log.Printf("[repository:administrator][GetAll] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:administrator][GetAll] error reading admins: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	log.Printf("[repository:administrator][GetAll] successfully fetched %d administrators", len(admins))
	return &admins, nil
}

func (r *AdministratorRepository) GetList(ctx context.Context) (*[]dto.AdministratorDTO, error) {
	var admins []dto.AdministratorDTO

	query := `
		SELECT id, first_name, last_name, email, gender, birth, phone
		FROM administrator
		WHERE deleted_at IS NULL
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:administrator][GetList] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("[repository:administrator][GetList] failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		var admin dto.AdministratorDTO

		err = rows.Scan(
			&admin.Id, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Gender, &admin.Birth, &admin.Phone,
		)
		if err != nil {
			log.Printf("[repository:administrator][GetList] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:administrator][GetList] error Reading admins: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	log.Printf("[repository:administrator][GetList] successfully fetched %d administrators", len(admins))
	return &admins, nil
}

func (r *AdministratorRepository) GetById(ctx context.Context, id uuid.UUID) (*dto.AdministratorDTO, error) {
	var admin dto.AdministratorDTO

	query := `
		SELECT id, first_name, last_name, email, gender, birth, phone
		FROM administrator
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&admin.Id, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Gender, &admin.Birth, &admin.Phone)
	if err != nil {
		log.Printf("[repository:administrator][GetById] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	log.Printf("[repository:administrator][GetById] successfully fetched administrator")
	return &admin, nil
}

func (r *AdministratorRepository) GetByEmail(ctx context.Context, email string) (*administrators.Administrator, error) {
	var (
		id                                              uuid.UUID
		firstName, lastName, emailStr, password, gender string
		lastLoginAt, createdAt, updatedAt               time.Time
		birth, deletedAt                                *time.Time
		phone                                           *string
	)

	query := `
		SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
		FROM administrator
		WHERE email = $1
	`

	err := r.Db.QueryRowContext(ctx, query, email).Scan(
		&id, &firstName, &lastName, &emailStr, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		log.Printf("[repository:administrator][GetByEmail] error executing SQL query '%s'\nfailed to fetch administrator: %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	admin := administrators.NewAdministratorFromDB(id, firstName, lastName, emailStr, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	log.Printf("[repository:administrator][GetByEmail] successfully fetched administrator")
	return admin, nil
}

func (r *AdministratorRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM administrator
			WHERE id = $1 
				AND deleted_at IS NULL)
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:administrator][ExistById] error executing SQL query '%s': %v", query, err)
		return false, err
	}

	log.Printf("[repository:administrator][ExistsById] id=%s exists=%t", id, exist)
	return exist, nil
}

func (r *AdministratorRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM administrator
			WHERE email = $1 
				AND deleted_at IS NULL)
	`
	err := r.Db.QueryRowContext(ctx, query, email).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:administrator][ExistByEmail] error executing SQL query '%s': %v", query, err)
		return false, err
	}

	log.Printf("[repository:administrator][ExistByEmail] email=%s exists=%t", email, exist)
	return exist, nil
}

func (r *AdministratorRepository) Create(ctx context.Context, adm *administrators.Administrator) (*administrators.Administrator, error) {
	var (
		id                                 uuid.UUID
		firstName, lastName, email, gender string
		lastLoginAt, createdAt, updatedAt  time.Time
		birth, birthVal, deletedAt         *time.Time
		phone, phoneVal                    *string
	)

	if adm.Birth() != nil {
		birthVal = adm.Birth().Value()
	}
	if adm.Phone() != nil {
		s := adm.Phone().String()
		phoneVal = s
	}

	query := `
		INSERT INTO administrator(id, first_name, last_name, email, password, gender, birth, phone)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, first_name, last_name, email, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(
		ctx, query, adm.Id(), adm.FirstName(), adm.LastName(), adm.Email().Value(), adm.Password().String(), adm.Gender(), birthVal, phoneVal).Scan(
		&id, &firstName, &lastName, &email, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Create] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	admin := administrators.NewAdministratorFromDB(id, firstName, lastName, email, "", gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	log.Printf("[repository:administrator][Create] successfully created administrator in DB %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) Update(ctx context.Context, adm *administrators.Administrator) (*administrators.Administrator, error) {
	var (
		id                                 uuid.UUID
		firstName, lastName, email, gender string
		phone                              *string
		lastLoginAt, createdAt, updatedAt  time.Time
		birth, deletedAt                   *time.Time
	)

	if adm.Birth() != nil {
		birth = adm.Birth().Value()
	}

	if adm.Phone() != nil {
		phone = adm.Phone().String()
	}

	query := `
        UPDATE administrator
        SET first_name = $1, last_name = $2, email = $3, password = $4, gender = $5, birth = $6, phone = $7, last_login_at = $8, updated_at = $9 
        WHERE id = $10
        RETURNING id, first_name, last_name, email, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
    `

	err := r.Db.QueryRowContext(
		ctx, query, adm.FirstName(), adm.LastName(), adm.Email().Value(), adm.Password().String(), adm.Gender(), birth, phone, adm.LastLoginAt, adm.UpdatedAt, adm.Id(),
	).Scan(
		&id, &firstName, &lastName, &email, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Update] error executing SQL query: %v", err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	updatedAdmin := administrators.NewAdministratorFromDB(
		id, firstName, lastName, email, adm.Password().String(), gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt,
	)

	return updatedAdmin, nil
}

func (r *AdministratorRepository) Delete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	var (
		firstName, lastName, email, gender string
		lastLoginAt, createdAt, updatedAt  time.Time
		birth, deletedAt                   *time.Time
		phone                              *string
	)

	query := `
		UPDATE administrator
		SET deleted_at = NOW() 
		WHERE id = $1 
		RETURNING first_name, last_name, email, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&firstName, &lastName, &email, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Delete] error executing SQL query '%s': %v", query, err)
		return nil, err
	}

	admin := administrators.NewAdministratorFromDB(id, firstName, lastName, email, "", gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	log.Printf("[repository:administrator][Delete] successfully soft deleted administrator %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) Restore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	var (
		idNew                              uuid.UUID
		firstName, lastName, email, gender string
		lastLoginAt, createdAt, updatedAt  time.Time
		birth, deletedAt                   *time.Time
		phone                              *string
	)

	query := `
		UPDATE administrator
		SET deleted_at = NULL
		WHERE id = $1
		RETURNING id, first_name, last_name, email, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&idNew, &firstName, &lastName, &email, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Restore] error executing SQL query '%s': %v", query, err)
		return nil, err
	}

	admin := administrators.NewAdministratorFromDB(id, firstName, lastName, email, "", gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	log.Printf("[repository:administrator][Delete] successfully restore administrator %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) CountAll(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM administrator
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountAll] error executing SQL query in CountAll: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *AdministratorRepository) CountActive(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM administrator 
		WHERE deleted_at IS NULL
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountActive] error executing SQL query in CountActive: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *AdministratorRepository) CountDeleted(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*)
		FROM administrator
		WHERE deleted_at IS NOT NULL
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountDeleted] error executing SQL query in CountDeleted: %v", err)
		return 0, err
	}
	return count, nil
}

func NewAdministratorRepository(db *sql.DB) administrators.AdministratorRepository {
	return &AdministratorRepository{Db: db}
}
