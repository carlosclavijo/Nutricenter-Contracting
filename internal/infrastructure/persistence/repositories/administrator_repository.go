package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
	"time"
)

type AdministratorRepository struct {
	Db *sql.DB
}

const (
	got                       = "%w: got %w"
	QueryGetAllAdministrators = `SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
								FROM administrator`
	QueryGetListAdministrators = `SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM administrator
									WHERE deleted_at IS NULL`
	QueryGetAdministratorById = `SELECT first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM administrator
									WHERE id = $1`
	QueryGetAdministratorByEmail = `SELECT id, first_name, last_name, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM administrator
									WHERE email = $1`
	QueryExistAdministratorById = `SELECT EXISTS(SELECT 1 
									FROM administrator
									WHERE id = $1 
									AND deleted_at IS NULL)`
	QueryExistAdministratorByEmail = `SELECT EXISTS(SELECT 1 
										FROM administrator
										WHERE email = $1 
										AND deleted_at IS NULL)`
	QueryCreateAdministrator = `INSERT INTO administrator(id, first_name, last_name, email, password, gender, birth, phone)
								VALUES($1, $2, $3, $4, $5, $6, $7, $8)
								RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryUpdateAdministrator = `UPDATE administrator
        						SET first_name = $1, last_name = $2, email = $3, password = $4, gender = $5, birth = $6, phone = $7, last_login_at = $8, updated_at = $9 
        						WHERE id = $10
        						RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryDeleteAdministrator = `UPDATE administrator
								SET deleted_at = NOW() 
								WHERE id = $1 
								RETURNING first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryRestoreAdministrator = `UPDATE administrator
									SET deleted_at = NULL
									WHERE id = $1
									RETURNING first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryCountAllAdministrators = `SELECT COUNT(*)
									FROM administrator`
	QueryCountActiveAdministrators = `SELECT COUNT(*)
										FROM administrator 
										WHERE deleted_at IS NULL`
	QueryCountDeletedAdministrators = `SELECT COUNT(*)
										FROM administrator 
										WHERE deleted_at IS NOT NULL`
)

var (
	ErrQueryAdministrator         = errors.New("query failed")
	ErrScanAdministrator          = errors.New("scan failed")
	ErrConcatenatingAdministrator = errors.New("error concatenating administrator values from DB")
	ErrIterationRowsAdministrator = errors.New("rows iteration error")
)

func (r *AdministratorRepository) GetAll(ctx context.Context) ([]*administrators.Administrator, error) {
	var (
		admins                                       []*administrators.Administrator
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	rows, err := r.Db.QueryContext(ctx, QueryGetAllAdministrators)
	if err != nil {
		log.Printf("[repository:administrator][GetAll] error executing SQL query '%s': %v", QueryGetAllAdministrators, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("[repository:administrator][GetAll] failed to close rows: %v", err)
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Printf("[repository:administrator][GetAll] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf(got, ErrScanAdministrator, err)
		}

		admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:adminstrator][GetAll] error concatenating administrator values from DB")
			return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
		}

		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:administrator][GetAll] error reading admins: %v", err)
		return nil, fmt.Errorf(got, ErrIterationRowsAdministrator, err)
	}

	log.Printf("[repository:administrator][GetAll] successfully fetched %d administrators", len(admins))
	return admins, nil
}

func (r *AdministratorRepository) GetList(ctx context.Context) ([]*administrators.Administrator, error) {
	var (
		admins                                       []*administrators.Administrator
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	rows, err := r.Db.QueryContext(ctx, QueryGetListAdministrators)
	if err != nil {
		log.Printf("[repository:administrator][GetList] error executing SQL query '%s': %v", QueryGetListAdministrators, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("[repository:administrator][GetList] failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(
			&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
		)
		if err != nil {
			log.Printf("[repository:administrator][GetList] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf(got, ErrScanAdministrator, err)
		}

		admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:adminstrator][GetList] error concatenating administrator values from DB")
			return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
		}

		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:administrator][GetList] error Reading admins: %v", err)
		return nil, fmt.Errorf(got, ErrIterationRowsAdministrator, err)
	}

	log.Printf("[repository:administrator][GetList] successfully fetched %d administrators", len(admins))
	return admins, nil
}

func (r *AdministratorRepository) GetById(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	err := r.Db.QueryRowContext(ctx, QueryGetAdministratorById, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		log.Printf("[repository:administrator][GetById] error executing SQL query '%s': %v", QueryGetAdministratorById, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][GetById] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][GetById] successfully fetched administrator")
	return admin, nil
}

func (r *AdministratorRepository) GetByEmail(ctx context.Context, email string) (*administrators.Administrator, error) {
	var (
		id                                       uuid.UUID
		firstName, lastName, password, gender    string
		lastLoginAt, createdAt, updatedAt, birth time.Time
		deletedAt                                *time.Time
		phone                                    *string
	)

	err := r.Db.QueryRowContext(ctx, QueryGetAdministratorByEmail, email).Scan(
		&id, &firstName, &lastName, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		log.Printf("[repository:administrator][GetByEmail] error executing SQL query '%s'\nfailed to fetch administrator: %v", QueryGetAdministratorByEmail, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][GetByEmail] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][GetByEmail] successfully fetched administrator")
	return admin, nil
}

func (r *AdministratorRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	err := r.Db.QueryRowContext(ctx, QueryExistAdministratorById, id).Scan(&exist)
	if err != nil {
		log.Printf("[repository:administrator][ExistById] error executing SQL query '%s': %v", QueryExistAdministratorById, err)
		return false, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	log.Printf("[repository:administrator][ExistsById] id=%s exists=%t", id, exist)
	return exist, nil
}

func (r *AdministratorRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	err := r.Db.QueryRowContext(ctx, QueryExistAdministratorByEmail, email).Scan(&exist)
	if err != nil {
		log.Printf("[repository:administrator][ExistByEmail] error executing SQL query '%s': %v", QueryExistAdministratorByEmail, err)
		return false, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	log.Printf("[repository:administrator][ExistByEmail] email=%s exists=%t", email, exist)
	return exist, nil
}

func (r *AdministratorRepository) Create(ctx context.Context, adm *administrators.Administrator) (*administrators.Administrator, error) {
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

	err := r.Db.QueryRowContext(
		ctx, QueryCreateAdministrator, adm.Id(), adm.FirstName(), adm.LastName(), adm.Email().Value(), adm.Password().String(), adm.Gender(), adm.Birth().Value(), phoneVal,
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Create] error executing SQL query '%s': %v", QueryCreateAdministrator, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][Create] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][Create] successfully created administrator in DB %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) Update(ctx context.Context, adm *administrators.Administrator) (*administrators.Administrator, error) {
	var (
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		phone                                        *string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
	)

	if adm.Phone() != nil {
		phone = adm.Phone().String()
	}

	err := r.Db.QueryRowContext(
		ctx, QueryUpdateAdministrator, adm.FirstName(), adm.LastName(), adm.Email().Value(), adm.Password().String(), adm.Gender(), adm.Birth().Value(), phone, adm.LastLoginAt(), adm.UpdatedAt(), adm.Id(),
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Update] error executing SQL query: %v", err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, adm.Password().String(), gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][Update] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][Update] successfully update administrator %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) Delete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	err := r.Db.QueryRowContext(ctx, QueryDeleteAdministrator, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:administrator][Delete] error executing SQL query '%s': %v", QueryDeleteAdministrator, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][Delete] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][Delete] successfully deleted administrator %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) Restore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	err := r.Db.QueryRowContext(ctx, QueryRestoreAdministrator, id).Scan(&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)

	if err != nil {
		log.Printf("[repository:administrator][Restore] error executing SQL query '%s': %v", QueryRestoreAdministrator, err)
		return nil, fmt.Errorf(got, ErrQueryAdministrator, err)
	}

	admin, err := administrators.NewAdministratorFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:adminstrator][Restore] error concatenating administrator values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingAdministrator, err)
	}

	log.Printf("[repository:administrator][Restore] successfully restore administrator %v", admin)
	return admin, nil
}

func (r *AdministratorRepository) CountAll(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountAllAdministrators).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountAll] error executing SQL query in CountAll: %v", err)
		return -1, fmt.Errorf(got, ErrQueryAdministrator, err)
	}
	return count, nil
}

func (r *AdministratorRepository) CountActive(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountActiveAdministrators).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountActive] error executing SQL query in CountActive: %v", err)
		return -1, fmt.Errorf(got, ErrQueryAdministrator, err)
	}
	return count, nil
}

func (r *AdministratorRepository) CountDeleted(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountDeletedAdministrators).Scan(&count)
	if err != nil {
		log.Printf("[repository:administrator][CountDeleted] error executing SQL query in CountDeleted: %v", err)
		return -1, fmt.Errorf(got, ErrQueryAdministrator, err)
	}
	return count, nil
}

func NewAdministratorRepository(db *sql.DB) administrators.AdministratorRepository {
	return &AdministratorRepository{Db: db}
}
