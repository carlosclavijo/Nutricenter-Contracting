package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

var (
	ErrDatabasePatient = errors.New("database is down")
)

func TestPatientRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()

	rows := sqlmock.NewRows(columns)
	for _, tc := range cases {
		rows.AddRow(
			tc.id, tc.firstName, tc.lastName, tc.email, tc.password, tc.gender,
			tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
		)
	}

	mock.ExpectQuery(QueryGetAllPatients).WillReturnRows(rows)

	ptnts, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, ptnts, len(cases))

	for i, tc := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			testCasesP(t, tc, ptnts[i])
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetAll_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	mock.ExpectQuery(QueryGetAllPatients).WillReturnError(ErrDatabasePatient)

	ptnts, err := repo.GetAll(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetAll_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	newColumns := []string{"id", "first_name"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New(), "John")

	mock.ExpectQuery(QueryGetAllPatients).WillReturnRows(rows)

	ptnts, err := repo.GetAll(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrScanPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetAll_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "Invalid", "User", "invalid-email", "$2a$10$abcdefghijklmnopqrstuvabcdefghijklmnopqrstuvab", "male", time.Now().AddDate(-10, 0, 0), nil, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetAllPatients).WillReturnRows(rows)

	ptnts, err := repo.GetAll(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()

	var filteredCases []Case
	for _, tc := range cases {
		if tc.deletedAt == nil {
			filteredCases = append(filteredCases, tc)
		}
	}

	filteredRows := sqlmock.NewRows(columns)
	for _, tc := range filteredCases {
		filteredRows.AddRow(
			tc.id, tc.firstName, tc.lastName, tc.email, tc.password, tc.gender,
			tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
		)
	}

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetListPatients)).WillReturnRows(filteredRows)

	ptnts, err := repo.GetList(context.Background())
	assert.NoError(t, err)
	assert.Len(t, ptnts, len(filteredCases))

	for i, patient := range ptnts {
		t.Run(filteredCases[i].name, func(t *testing.T) {
			tc := filteredCases[i]
			testCasesP(t, tc, patient)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetList_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	mock.ExpectQuery(QueryGetListPatients).WillReturnError(ErrDatabasePatient)

	ptnts, err := repo.GetList(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetList_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	newColumns := []string{"id", "first_name"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New(), "John")

	mock.ExpectQuery(QueryGetListPatients).WillReturnRows(rows)

	ptnts, err := repo.GetList(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrScanPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetList_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "Invalid", "User", "invalid-email", "$2a$10$abcdefghijklmnopqrstuvabcdefghijklmnopqrstuvab", "male", time.Now().AddDate(-10, 0, 0), nil, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetListPatients).WillReturnRows(rows)

	ptnts, err := repo.GetList(context.Background())

	assert.Nil(t, ptnts)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()
	tc := cases[3]

	mockRows := sqlmock.NewRows([]string{
		"first_name", "last_name", "email", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		tc.firstName, tc.lastName, tc.email, tc.password, tc.gender,
		tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientById)).WithArgs(tc.id).WillReturnRows(mockRows)

	patient, err := repo.GetById(context.Background(), tc.id)
	assert.NoError(t, err)
	testCasesP(t, tc, patient)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPatientRepository_GetById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientById)).WithArgs(id).WillReturnError(ErrDatabasePatient)

	patient, err := repo.GetById(context.Background(), id)

	assert.Nil(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetById_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	id := uuid.New()
	invalidEmail := "not-an-email"

	mockRows := sqlmock.NewRows([]string{
		"first_name", "last_name", "email", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		"John", "Doe", invalidEmail, "$2a$10$dummyhash", "male",
		time.Now(), nil, time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientById)).WithArgs(id).WillReturnRows(mockRows)

	patient, err := repo.GetById(context.Background(), id)

	assert.Nil(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()
	tc := cases[3]

	mockRows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		tc.id, tc.firstName, tc.lastName, tc.password, tc.gender,
		tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientByEmail)).WithArgs(tc.email).WillReturnRows(mockRows)

	patient, err := repo.GetByEmail(context.Background(), tc.email)
	testCasesP(t, tc, patient)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPatientRepository_GetByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientByEmail)).WithArgs(email).WillReturnError(ErrDatabasePatient)

	patient, err := repo.GetByEmail(context.Background(), email)

	assert.Nil(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_GetByEmail_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email := "valid@email.com"
	password := "soft"

	mockRows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		uuid.New(), "John", "Doe", password, "male",
		time.Now(), nil, time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetPatientByEmail)).WithArgs(email).WillReturnRows(mockRows)

	patient, err := repo.GetByEmail(context.Background(), email)

	assert.Nil(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_ExistById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exists := tc.deletedAt == nil
			mockRow := sqlmock.NewRows([]string{"exists"}).AddRow(exists)

			mock.ExpectQuery(regexp.QuoteMeta(QueryExistPatientById)).WithArgs(tc.id).WillReturnRows(mockRow)

			result, err := repo.ExistById(context.Background(), tc.id)

			assert.NoError(t, err)
			assert.Equal(t, exists, result)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_ExistById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistPatientById)).WithArgs(id).WillReturnError(ErrDatabasePatient)

	patient, err := repo.ExistById(context.Background(), id)

	assert.False(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_ExistByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	cases := Cases()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exists := tc.deletedAt == nil
			mockRow := sqlmock.NewRows([]string{"exists"}).AddRow(exists)

			mock.ExpectQuery(regexp.QuoteMeta(QueryExistPatientByEmail)).WithArgs(tc.email).WillReturnRows(mockRow)

			result, err := repo.ExistByEmail(context.Background(), tc.email)

			assert.NoError(t, err)
			assert.Equal(t, exists, result)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_ExistByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistPatientByEmail)).WithArgs(email).WillReturnError(ErrDatabasePatient)

	patient, err := repo.ExistByEmail(context.Background(), email)

	assert.False(t, patient)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		ptn.Id(), ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
		ptn.Password().String(), ptn.Gender().String(),
		ptn.Birth().Value(), ptn.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreatePatient)).
		WithArgs(
			ptn.Id(), ptn.FirstName(), ptn.LastName(),
			ptn.Email().Value(), ptn.Password().String(),
			string(ptn.Gender()),
			ptn.Birth().Value(),
			ptn.Phone(),
		).WillReturnRows(rows)

	result, err := repo.Create(context.Background(), ptn)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ptn.Id(), result.Id())
	assert.Equal(t, ptn.FirstName(), result.FirstName())
	assert.Equal(t, ptn.LastName(), result.LastName())
	assert.Equal(t, ptn.Email(), result.Email())
	assert.Equal(t, ptn.Password(), result.Password())
	assert.Equal(t, ptn.Gender(), result.Gender())
	assert.Equal(t, ptn.Birth(), result.Birth())
	assert.Equal(t, ptn.Phone(), result.Phone())
	assert.NotNil(t, result.LastLoginAt())
	assert.NotNil(t, result.CreatedAt())
	assert.NotNil(t, result.UpdatedAt())
	assert.Nil(t, result.DeletedAt())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func TestPatientRepository_Create_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email, password, gender, birth := valueObjects(t, "test@example.com", "Str0ng!!1", "M", time.Now().AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreatePatient)).
		WithArgs(
			ptn.Id(), ptn.FirstName(), ptn.LastName(),
			ptn.Email().Value(), ptn.Password().String(),
			ptn.Gender(),
			ptn.Birth().Value(),
			ptn.Phone(),
		).
		WillReturnError(ErrDatabasePatient)

	result, err := repo.Create(context.Background(), ptn)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryPatient)
	assert.ErrorIs(t, err, ErrDatabasePatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Create_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email, password, gender, birth := valueObjects(t, "test@example.com", "Str0ng!!1", "M", time.Now().AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		ptn.Id(), ptn.FirstName(), ptn.LastName(),
		ptn.Email().Value(), ptn.Password().String(),
		"X",
		ptn.Birth().Value(),
		nil,
		time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreatePatient)).
		WithArgs(
			ptn.Id(), ptn.FirstName(), ptn.LastName(),
			ptn.Email().Value(), ptn.Password().String(),
			ptn.Gender(),
			ptn.Birth().Value(),
			ptn.Phone(),
		).
		WillReturnRows(rows)

	result, err := repo.Create(context.Background(), ptn)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingPatient)
	assert.Contains(t, err.Error(), "invalid")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		ptn.Id(), ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
		ptn.Password().String(), ptn.Gender().String(),
		ptn.Birth().Value(), ptn.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdatePatient)).
		WithArgs(
			ptn.FirstName(), ptn.LastName(), ptn.Email().Value(), ptn.Password().String(),
			string(ptn.Gender()), ptn.Birth().Value(), ptn.Phone(), ptn.LastLoginAt(), ptn.UpdatedAt(), ptn.Id(),
		).WillReturnRows(rows)

	result, err := repo.Update(context.Background(), ptn)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ptn.Id(), result.Id())
	assert.Equal(t, ptn.FirstName(), result.FirstName())
	assert.Equal(t, ptn.LastName(), result.LastName())
	assert.Equal(t, ptn.Email().Value(), result.Email().Value())
	assert.Equal(t, ptn.Password().String(), result.Password().String())
	assert.Equal(t, ptn.Gender(), result.Gender())
	assert.Equal(t, ptn.Birth(), result.Birth())
	assert.Equal(t, ptn.Phone(), result.Phone())
	assert.NotNil(t, result.LastLoginAt())
	assert.NotNil(t, result.CreatedAt())
	assert.NotNil(t, result.UpdatedAt())
	assert.Nil(t, result.DeletedAt())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Update_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", time.Now().AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdatePatient)).
		WithArgs(
			ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
			ptn.Password().String(), string(ptn.Gender()), ptn.Birth().Value(),
			ptn.Phone(), ptn.LastLoginAt(), ptn.UpdatedAt(), ptn.Id(),
		).
		WillReturnError(fmt.Errorf("database is down"))

	result, err := repo.Update(context.Background(), ptn)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database is down")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Update_NewPatientError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		uuid.Nil, ptn.FirstName(), ptn.LastName(), "",
		ptn.Password().String(), ptn.Gender().String(),
		ptn.Birth().Value(), ptn.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdatePatient)).
		WithArgs(
			ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
			ptn.Password().String(), string(ptn.Gender()), ptn.Birth().Value(),
			ptn.Phone(), ptn.LastLoginAt(), ptn.UpdatedAt(), ptn.Id(),
		).
		WillReturnRows(rows)

	result, err := repo.Update(context.Background(), ptn)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "concatenating patient values")
	assert.ErrorIs(t, err, ErrConcatenatingPatient)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPatientRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)
	nColumns := []string{
		"first_name", "last_name", "email", "password", "gender",
		"birth", "phone", "last_login_at", "created_at", "updated_at", "deleted_at",
	}
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(nColumns).AddRow(
			ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
			ptn.Password().String(), ptn.Gender().String(), ptn.Birth().Value(),
			ptn.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryDeletePatient)).
			WithArgs(ptn.Id()).
			WillReturnRows(rows)

		result, err := repo.Delete(context.Background(), ptn.Id())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, ptn.FirstName(), result.FirstName())
		assert.Equal(t, ptn.LastName(), result.LastName())
		assert.Equal(t, ptn.Email().Value(), result.Email().Value())
		assert.Equal(t, ptn.Password().String(), result.Password().String())
		assert.Equal(t, ptn.Gender(), result.Gender())
		assert.Equal(t, ptn.Birth(), result.Birth())
		assert.Equal(t, ptn.Phone(), result.Phone())
		assert.NotNil(t, result.DeletedAt())

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryDeletePatient)).
			WithArgs(ptn.Id()).
			WillReturnError(ErrDatabasePatient)

		result, err := repo.Delete(context.Background(), ptn.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryPatient)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("concatenation error", func(t *testing.T) {
		invalidEmail := ""
		rows := sqlmock.NewRows(nColumns).AddRow(
			ptn.FirstName(), ptn.LastName(), invalidEmail,
			ptn.Password().String(), ptn.Gender().String(), ptn.Birth().Value(),
			ptn.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryDeletePatient)).
			WithArgs(ptn.Id()).
			WillReturnRows(rows)

		result, err := repo.Delete(context.Background(), ptn.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConcatenatingPatient)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestPatientRepository_Restore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	ptn := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)
	nColumns := []string{
		"first_name", "last_name", "email", "password", "gender",
		"birth", "phone", "last_login_at", "created_at", "updated_at", "deleted_at",
	}
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(nColumns).AddRow(
			ptn.FirstName(), ptn.LastName(), ptn.Email().Value(),
			ptn.Password().String(), ptn.Gender().String(), ptn.Birth().Value(),
			ptn.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryRestorePatient)).
			WithArgs(ptn.Id()).
			WillReturnRows(rows)

		result, err := repo.Restore(context.Background(), ptn.Id())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, ptn.FirstName(), result.FirstName())
		assert.Equal(t, ptn.LastName(), result.LastName())
		assert.Equal(t, ptn.Email().Value(), result.Email().Value())
		assert.Equal(t, ptn.Password().String(), result.Password().String())
		assert.Equal(t, ptn.Gender(), result.Gender())
		assert.Equal(t, ptn.Birth(), result.Birth())
		assert.Equal(t, ptn.Phone(), result.Phone())
		assert.NotNil(t, result.DeletedAt())

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryRestorePatient)).
			WithArgs(ptn.Id()).
			WillReturnError(ErrDatabasePatient)

		result, err := repo.Restore(context.Background(), ptn.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryPatient)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("concatenation error", func(t *testing.T) {
		invalidEmail := ""
		rows := sqlmock.NewRows(nColumns).AddRow(
			ptn.FirstName(), ptn.LastName(), invalidEmail,
			ptn.Password().String(), ptn.Gender().String(), ptn.Birth().Value(),
			ptn.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryRestorePatient)).
			WithArgs(ptn.Id()).
			WillReturnRows(rows)

		result, err := repo.Restore(context.Background(), ptn.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConcatenatingPatient)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestPatientRepository_CountAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountAllPatients)).
			WillReturnRows(rows)

		count, err := repo.CountAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountAllPatients)).
			WillReturnError(ErrDatabasePatient)

		count, err := repo.CountAll(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryPatient)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestPatientRepository_CountActive(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountActivePatients)).WillReturnRows(rows)
		count, err := repo.CountActive(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountActivePatients)).
			WillReturnError(ErrDatabasePatient)

		count, err := repo.CountActive(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryPatient)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestPatientRepository_CountDeleted(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPatientRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountDeletedPatients)).WillReturnRows(rows)
		count, err := repo.CountDeleted(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountDeletedPatients)).WillReturnError(ErrDatabasePatient)
		count, err := repo.CountDeleted(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryPatient)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func testCasesP(t *testing.T, tc Case, patient *patients.Patient) {
	assert.NotNil(t, patient)
	assert.Equal(t, tc.id, patient.Id())
	assert.Equal(t, tc.firstName, patient.FirstName())
	assert.Equal(t, tc.lastName, patient.LastName())
	assert.Equal(t, tc.email, patient.Email().Value())
	assert.Equal(t, tc.password, patient.Password().String())
	assert.Equal(t, tc.gender, patient.Gender().String())
	assert.True(t, tc.birth.Equal(patient.Birth().Value()))
	if tc.phone != nil {
		phone, _ := vo.NewPhone(tc.phone)
		assert.NotNil(t, patient.Phone())
		assert.Equal(t, phone, patient.Phone())
	} else {
		assert.Nil(t, patient.Phone())
	}
	assert.True(t, tc.lastLoginAt.Equal(patient.LastLoginAt()))
	assert.True(t, tc.createdAt.Equal(patient.CreatedAt()))
	assert.True(t, tc.updatedAt.Equal(patient.UpdatedAt()))
	if tc.deletedAt != nil {
		assert.NotNil(t, patient.DeletedAt())
		assert.True(t, tc.deletedAt.Equal(*patient.DeletedAt()))
	} else {
		assert.Nil(t, patient.DeletedAt())
	}
}
