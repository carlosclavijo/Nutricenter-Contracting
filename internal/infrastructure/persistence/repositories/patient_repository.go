package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/google/uuid"
	"log"
	"time"
)

type PatientRepository struct {
	Db *sql.DB
}

const (
	QueryGetAllPatients = `SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM patient`
	QueryGetListPatients = `SELECT id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at 
									FROM patient 
									WHERE deleted_at IS NULL`
	QueryGetPatientById = `SELECT first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM patient
									WHERE id = $1`
	QueryGetPatientByEmail = `SELECT id, first_name, last_name, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at
									FROM patient
									WHERE email = $1`
	QueryExistPatientById = `SELECT EXISTS(SELECT 1 
									FROM patient
									WHERE id = $1 
									AND deleted_at IS NULL)`
	QueryExistPatientByEmail = `SELECT EXISTS(SELECT 1 
										FROM patient
										WHERE email = $1 
										AND deleted_at IS NULL)`
	QueryCreatePatient = `INSERT INTO patient(id, first_name, last_name, email, password, gender, birth, phone)
								VALUES($1, $2, $3, $4, $5, $6, $7, $8)
								RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryUpdatePatient = `UPDATE patient
        						SET first_name = $1, last_name = $2, email = $3, password = $4, gender = $5, birth = $6, phone = $7, last_login_at = $8, updated_at = $9 
        						WHERE id = $10
        						RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryDeletePatient = `UPDATE patient
								SET deleted_at = NOW() 
								WHERE id = $1 
								RETURNING first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryRestorePatient = `UPDATE patient
									SET deleted_at = NULL
									WHERE id = $1
									RETURNING id, first_name, last_name, email, password, gender, birth, phone, last_login_at, created_at, updated_at, deleted_at`
	QueryCountAllPatients = `SELECT COUNT(*)
									FROM patient`
	QueryCountActivePatients = `SELECT COUNT(*)
										FROM patient 
										WHERE deleted_at IS NULL`
	QueryCountDeletedPatients = `SELECT COUNT(*)
										FROM patient 
										WHERE deleted_at IS NOT NULL`
)

var (
	ErrQueryPatient         = errors.New("query failed")
	ErrScanPatient          = errors.New("scan failed")
	ErrConcatenatingPatient = errors.New("error concatenating patient values from DB")
	ErrIterationRowsPatient = errors.New("rows iteration error")
)

func (r *PatientRepository) GetAll(ctx context.Context) ([]*patients.Patient, error) {
	var (
		ptns                                         []*patients.Patient
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	rows, err := r.Db.QueryContext(ctx, QueryGetAllPatients)
	if err != nil {
		log.Printf("[repository:patient][GetAll] error executing SQL query '%s': %v", QueryGetAllPatients, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
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
			return nil, fmt.Errorf(got, ErrScanPatient, err)
		}

		patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:patient][GetAll] error concatenating patient values from DB")
			return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
		}

		ptns = append(ptns, patient)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:patient][GetAll] error reading patients: %v", err)
		return nil, fmt.Errorf(got, ErrIterationRowsPatient, err)
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

	rows, err := r.Db.QueryContext(ctx, QueryGetListPatients)
	if err != nil {
		log.Printf("[repository:patient][GetList] error executing SQL query '%s': %v", QueryGetListPatients, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("[repository:patient][GetList] failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(
			&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
		)
		if err != nil {
			log.Printf("[repository:patient][GetList] error reading patientDTO for a slice of patients: %v", err)
			return nil, fmt.Errorf(got, ErrScanPatient, err)
		}

		patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			log.Printf("[repository:patient][GetList] error concatenating patient values from DB")
			return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
		}

		ptnts = append(ptnts, patient)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:patient][GetList] error Reading patients: %v", err)
		return nil, fmt.Errorf(got, ErrIterationRowsPatient, err)
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

	err := r.Db.QueryRowContext(ctx, QueryGetPatientById, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][GetById] error executing SQL query '%s': %v", QueryGetPatientById, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	ptnt, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][GetById] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
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

	err := r.Db.QueryRowContext(ctx, QueryGetPatientByEmail, email).Scan(
		&id, &firstName, &lastName, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		log.Printf("[repository:patient][GetByEmail] error executing SQL query '%s'\nfailed to fetch patient: %v", QueryGetPatientByEmail, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][GetByEmail] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
	}

	log.Printf("[repository:patient][GetByEmail] successfully fetched patient")
	return patient, nil
}

func (r *PatientRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	err := r.Db.QueryRowContext(ctx, QueryExistPatientById, id).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:patient][ExistById] error executing SQL query '%s': %v", QueryExistPatientById, err)
		return false, fmt.Errorf(got, ErrQueryPatient, err)
	}

	log.Printf("[repository:patient][ExistsById] id=%s exists=%t", id, exist)
	return exist, nil
}

func (r *PatientRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	err := r.Db.QueryRowContext(ctx, QueryExistPatientByEmail, email).Scan(
		&exist,
	)
	if err != nil {
		log.Printf("[repository:patient][ExistByEmail] error executing SQL query '%s': %v", QueryExistPatientByEmail, err)
		return false, fmt.Errorf(got, ErrQueryPatient, err)
	}

	log.Printf("[repository:patient][ExistByEmail] email=%s exists=%t", email, exist)
	return exist, nil
}

func (r *PatientRepository) Create(ctx context.Context, ptn *patients.Patient) (*patients.Patient, error) {
	var (
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone, phoneVal                              *string
	)

	if ptn.Phone() != nil {
		s := ptn.Phone().String()
		phoneVal = s
	}

	err := r.Db.QueryRowContext(
		ctx, QueryCreatePatient, ptn.Id(), ptn.FirstName(), ptn.LastName(), ptn.Email().Value(), ptn.Password().String(), ptn.Gender(), ptn.Birth().Value(), phoneVal,
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Create] error executing SQL query '%s': %v", QueryCreatePatient, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][Create] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
	}

	log.Printf("[repository:patient][Create] successfully created patient in DB %v", patient)
	return patient, nil
}

func (r *PatientRepository) Update(ctx context.Context, ptn *patients.Patient) (*patients.Patient, error) {
	var (
		id                                           uuid.UUID
		firstName, lastName, email, password, gender string
		phone                                        *string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
	)

	if ptn.Phone() != nil {
		phone = ptn.Phone().String()
	}

	err := r.Db.QueryRowContext(
		ctx, QueryUpdatePatient, ptn.FirstName(), ptn.LastName(), ptn.Email().Value(), ptn.Password().String(), ptn.Gender(), ptn.Birth().Value(), phone, ptn.LastLoginAt(), ptn.UpdatedAt(), ptn.Id(),
	).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Update] error executing SQL query: %v", err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:patient][Update] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
	}

	return patient, nil
}

func (r *PatientRepository) Delete(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	err := r.Db.QueryRowContext(ctx, QueryDeletePatient, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Delete] error executing SQL query '%s': %v", QueryDeletePatient, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:admiinstrator][Delete] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
	}

	log.Printf("[repository:patient][Delete] successfully soft deleted patient %v", patient)
	return patient, nil
}

func (r *PatientRepository) Restore(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	var (
		firstName, lastName, email, password, gender string
		lastLoginAt, createdAt, updatedAt, birth     time.Time
		deletedAt                                    *time.Time
		phone                                        *string
	)

	err := r.Db.QueryRowContext(ctx, QueryRestorePatient, id).Scan(
		&firstName, &lastName, &email, &password, &gender, &birth, &phone, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:patient][Restore] error executing SQL query '%s': %v", QueryRestorePatient, err)
		return nil, fmt.Errorf(got, ErrQueryPatient, err)
	}

	patient, err := patients.NewPatientFromDB(id, firstName, lastName, email, password, gender, birth, phone, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		log.Printf("[repository:admiinstrator][Restore] error concatenating patient values from DB")
		return nil, fmt.Errorf(got, ErrConcatenatingPatient, err)
	}

	log.Printf("[repository:patient][Delete] successfully restore patient %v", patient)
	return patient, nil
}

func (r *PatientRepository) CountAll(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountAllPatients).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountAll] error executing SQL query in CountAll: %v", err)
		return -1, fmt.Errorf(got, ErrQueryPatient, err)
	}
	return count, nil
}

func (r *PatientRepository) CountActive(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountActivePatients).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountActive] error executing SQL query in CountActive: %v", err)
		return -1, fmt.Errorf(got, ErrQueryPatient, err)
	}
	return count, nil
}

func (r *PatientRepository) CountDeleted(ctx context.Context) (int, error) {
	var count int

	err := r.Db.QueryRowContext(ctx, QueryCountDeletedPatients).Scan(&count)
	if err != nil {
		log.Printf("[repository:patient][CountDeleted] error executing SQL query in CountDeleted: %v", err)
		return -1, fmt.Errorf(got, ErrQueryPatient, err)
	}
	return count, nil
}

func NewPatientRepository(db *sql.DB) patients.PatientRepository {
	return &PatientRepository{Db: db}
}
