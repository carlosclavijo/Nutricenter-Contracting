package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/google/uuid"
	"log"
	"time"
)

type PatientRepository struct {
	Db *sql.DB
}

func (r *PatientRepository) GetAll(ctx context.Context) ([]*patients.Patient, error) {
	var (
		ptns                                         []*patients.Patient
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	query := `
		SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
		FROM patient
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:patient][GetAll] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Printf("[repository:patient][GetAll] failed to close rows: %v", err)
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(
			&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
		)
		if err != nil {
			log.Printf("[repository:patient][GetAll] error reading patientDTO for a slice of patients: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:patient][GetAll] error concatenating patient values from DB")
			return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
		}

		ptns = append(ptns, patient)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:patient][GetAll] error reading patients: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	log.Printf("[repository:patient][GetAll] successfully fetched %d patients", len(ptns))
	return ptns, nil
}

func (r *PatientRepository) GetList(ctx context.Context) ([]*patients.Patient, error) {
	var (
		ptnts                                        []*patients.Patient
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	query := `
		SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
		FROM patient
		WHERE deleted_at IS NULL
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:patient][GetList] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("[repository:patient][GetList] failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(
			id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt,
		)
		if err != nil {
			log.Printf("[repository:patient][GetList] error reading patientDTO for a slice of patients: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:patient][GetList] error concatenating patient values from DB")
			return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
		}

		ptnts = append(ptnts, patient)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:patient][GetList] error Reading patients: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	log.Printf("[repository:patient][GetList] successfully fetched %d patients", len(ptnts))
	return ptnts, nil
}

func (r *PatientRepository) GetById(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	query := `
		SELECT first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
		FROM patient
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][GetById] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	ptnt, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][GetById] error concatenating patient values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:patient][GetById] successfully fetched patient")
	return ptnt, nil
}

func (r *PatientRepository) GetByEmail(ctx context.Context, email string) (*patients.Patient, error) {
	var (
		id                                       uuid.UUID
		firstName, lastName, password, gender    string
		lastLoginAt, createdAt, updatedAt, birth time.Time
		deletedAt                                *time.Time
		phone                                    *string
	)

	query := `
		SELECT id, first_name, last_name, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
		FROM patient
		WHERE email = $1
	`

	err := r.Db.QueryRowContext(ctx, query, email).Scan(
		&id, &firstName, &lastName, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		log.Printf("[repository:patient][GetByEmail] error executing SQL query '%s'\nfailed to fetch patient: %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][GetByEmail] error concatenating patient values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:patient][GetByEmail] successfully fetched patient")
	return patient, nil
}

func (r *PatientRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM patient
			WHERE id = $1 
				AND deleted_at IS NULL)
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:patient][ExistById] error executing SQL query '%s': %v", query, err)
		return false, err
	}

	log.Printf("[repository:patient][ExistsById] id=%s exists=%t", id, exist)
	return exist, nil
}

func (r *PatientRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM patient
			WHERE email = $1 
				AND deleted_at IS NULL)
	`
	err := r.Db.QueryRowContext(ctx, query, email).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:patient][ExistByEmail] error executing SQL query '%s': %v", query, err)
		return false, err
	}

	log.Printf("[repository:patient][ExistByEmail] email=%s exists=%t", email, exist)
	return exist, nil
}

func (r *PatientRepository) Create(ctx context.Context, adm *patients.Patient) (*patients.Patient, error) {
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
		INSERT INTO patient(id, first_name, last_name, email, password, gender, birth, phone)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(
		ctx, query, adm.Id, adm.FirstName, adm.LastName, adm.Email().Value(), adm.Password().String(), adm.Gender, adm.Birth, phoneVal,
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Create] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][Create] error concatenating patient values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:patient][Create] successfully created patient in DB %v", patient)
	return patient, nil
}

func (r *PatientRepository) Update(ctx context.Context, adm *patients.Patient) (*patients.Patient, error) {
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

	query := `
        UPDATE patient
        SET first_name = $1, last_name = $2, email = $3, password = $4, gender = $5, birth = $6, phone = $7, last_login_at = $8, updated_at = $9 
        WHERE id = $10
        RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
    `

	err := r.Db.QueryRowContext(
		ctx, query, adm.FirstName, adm.LastName, adm.Email().Value(), adm.Password().String(), adm.Gender, birth, phone, adm.LastLoginAt, adm.UpdatedAt, adm.Id(),
	).Scan(
		&id, &firstName, &lastName, &email, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Update] error executing SQL query: %v", err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	admin, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][Update] error concatenating patient values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	return admin, nil
}

func (r *PatientRepository) Delete(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	var (
		idNew                                        uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	query := `
		UPDATE patient
		SET deleted_at = NOW() 
		WHERE id = $1 
		RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&idNew, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Delete] error executing SQL query '%s': %v", query, err)
		return nil, err
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:admiinstrator][Delete] error concatenating administrator values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:patient][Delete] successfully soft deleted patient %v", patient)
	return patient, nil
}

func (r *PatientRepository) Restore(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	var (
		idNew                                        uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	query := `
		UPDATE patient
		SET deleted_at = NULL
		WHERE id = $1
		RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(
		&idNew, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Restore] error executing SQL query '%s': %v", query, err)
		return nil, err
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:admiinstrator][Restore] error concatenating administrator values from DB")
		return nil, fmt.Errorf("%w: error concatenating administrator values from DB", err)
	}

	log.Printf("[repository:patient][Delete] successfully restore patient %v", patient)
	return patient, nil
}

func (r *PatientRepository) CountAll(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM patient
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountAll] error executing SQL query in CountAll: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *PatientRepository) CountActive(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*)
		FROM patient 
		WHERE deleted_at IS NULL
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountActive] error executing SQL query in CountActive: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *PatientRepository) CountDeleted(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*)
		FROM patient
		WHERE deleted_at IS NOT NULL
	`
	err := r.Db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountDeleted] error executing SQL query in CountDeleted: %v", err)
		return 0, err
	}
	return count, nil
}

func NewPatientRepository(db *sql.DB) patients.PatientRepository {
	return &PatientRepository{Db: db}
}
