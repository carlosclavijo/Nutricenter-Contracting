package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

var (
	columns                  = []string{"id", "first_name", "last_name", "email", "password", "gender", "birth", "phone", "last_login_at", "created_at", "updated_at", "deleted_at"}
	ErrDatabaseAdministrator = errors.New("database is down")
)

type Case struct {
	name                                         string
	id                                           uuid.UUID
	firstName, lastName, email, password, gender string
	birth                                        time.Time
	phone                                        *string
	lastLoginAt, createdAt, updatedAt            time.Time
	deletedAt                                    *time.Time
}

func Cases() []Case {
	phones := []string{"78787878", "7744665577", "67829348", "70926048", "77141516"}
	now := time.Now()
	deletedAtDates := []time.Time{now.AddDate(0, 0, -5), now.AddDate(0, 0, -15), now.AddDate(0, -1, 0)}
	return []Case{
		{"Case 1", uuid.New(), "John", "Doe", "john@doe.com", "$2a$10$A1b2C3d4E5f6G7h8I9j0K1l2M3n4O5p6Q7r8S9t0U1v2W3x4Y5z6A", "male", now.AddDate(-20, 0, 0), &phones[0], now, now, now, &deletedAtDates[0]},
		{"Case 2", uuid.New(), "Jane", "Doe", "jane@doe.com", "$2a$10$Z9y8X7w6V5u4T3s2R1q0P9o8N7m6L5k4J3i2H1g0F9e8D7c6B5a4A", "female", now.AddDate(-19, 0, 0), &phones[1], now, now, now, nil},
		{"Case 3", uuid.New(), "Alice", "Smith", "alice@smith.com", "$2a$10$q1W2e3R4t5Y6u7I8o9P0a1S2d3F4g5H6j7K8l9Z0x1C2v3B4n5M6A", "female", now.AddDate(-25, 0, 0), &phones[2], now, now, now, nil},
		{"Case 4", uuid.New(), "Bob", "Smith", "bob@smith.com", "$2a$10$P0o9I8u7Y6t5R4e3W2q1A0s9D8f7G6h5J4k3L2z1X0c9V8b7N6m5A", "undefined", now.AddDate(-30, 0, 0), &phones[3], now, now, now, nil},
		{"Case 5", uuid.New(), "Carlos", "Gomez", "carlos@gomez.com", "$2a$10$m1N2b3V4c5X6z7L8k9J0h1G2f3D4s5A6q7W8e9R0t1Y2u3I4o5PA6", "male", now.AddDate(-22, 0, 0), &phones[4], now, now, now, &deletedAtDates[1]},
		{"Case 6", uuid.New(), "Maria", "Gomez", "maria@gomez.com", "$2a$10$Q1w2E3r4T5y6U7i8O9p0A1s2D3f4G5h6J7k8L9z0X1c2V3b4N5m6A", "female", now.AddDate(-28, 0, 0), nil, now, now, now, nil},
		{"Case 7", uuid.New(), "David", "Lee", "david@lee.com", "$2a$10$N6m5B4v3C2x1Z0l9K8j7H6g5F4d3S2a1Q0w9E8r7T6y5U4i3O2p1A", "male", now.AddDate(-35, 0, 0), nil, now, now, now, nil},
		{"Case 8", uuid.New(), "Sophia", "Lee", "sophia@lee.com", "$2a$10$O9p8I7u6Y5t4R3e2W1q0A9s8D7f6G5h4J3k2L1z0X9c8V7b6N5m4A", "undefined", now.AddDate(-27, 0, 0), nil, now, now, now, &deletedAtDates[2]},
		{"Case 9", uuid.New(), "Lucas", "Martinez", "lucas@martinez.com", "$2a$10$L1k2J3h4G5f6D7s8A9q0W1e2R3t4Y5u6I7o8P9z0X1c2V3b4N5m6A", "male", now.AddDate(-21, 0, 0), nil, now, now, now, nil},
		{"Case 10", uuid.New(), "Emma", "Martinez", "emma@martinez.com", "$2a$10$Z0x9C8v7B6n5M4l3K2j1H0g9F8d7S6a5Q4w3E2r1T0y9U8i7O6p5A", "female", now.AddDate(-26, 0, 0), nil, now, now, now, nil},
	}
}

func TestAdministratorRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	cases := Cases()

	rows := sqlmock.NewRows(columns)
	for _, tc := range cases {
		rows.AddRow(
			tc.id, tc.firstName, tc.lastName, tc.email, tc.password, tc.gender,
			tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
		)
	}

	mock.ExpectQuery(QueryGetAllAdministrators).WillReturnRows(rows)

	admins, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, admins, len(cases))

	for i, tc := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			testCases(t, tc, admins[i])
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetAll_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	mock.ExpectQuery(QueryGetAllAdministrators).WillReturnError(ErrDatabaseAdministrator)

	admins, err := repo.GetAll(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetAll_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	newColumns := []string{"id", "first_name"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New(), "John")

	mock.ExpectQuery(QueryGetAllAdministrators).WillReturnRows(rows)

	admins, err := repo.GetAll(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrScanAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetAll_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "Invalid", "User", "invalid-email", "$2a$10$abcdefghijklmnopqrstuvabcdefghijklmnopqrstuvab", "male", time.Now().AddDate(-10, 0, 0), nil, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetAllAdministrators).WillReturnRows(rows)

	admins, err := repo.GetAll(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
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

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetListAdministrators)).WillReturnRows(filteredRows)

	admins, err := repo.GetList(context.Background())
	assert.NoError(t, err)
	assert.Len(t, admins, len(filteredCases))

	for i, admin := range admins {
		t.Run(filteredCases[i].name, func(t *testing.T) {
			tc := filteredCases[i]
			testCases(t, tc, admin)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetList_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	mock.ExpectQuery(QueryGetListAdministrators).WillReturnError(ErrDatabaseAdministrator)

	admins, err := repo.GetList(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetList_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	newColumns := []string{"id", "first_name"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New(), "John")

	mock.ExpectQuery(QueryGetListAdministrators).WillReturnRows(rows)

	admins, err := repo.GetList(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrScanAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetList_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "Invalid", "User", "invalid-email", "$2a$10$abcdefghijklmnopqrstuvabcdefghijklmnopqrstuvab", "male", time.Now().AddDate(-10, 0, 0), nil, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetListAdministrators).WillReturnRows(rows)

	admins, err := repo.GetList(context.Background())

	assert.Nil(t, admins)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	cases := Cases()
	tc := cases[3]

	mockRows := sqlmock.NewRows([]string{
		"first_name", "last_name", "email", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		tc.firstName, tc.lastName, tc.email, tc.password, tc.gender,
		tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorById)).WithArgs(tc.id).WillReturnRows(mockRows)

	admin, err := repo.GetById(context.Background(), tc.id)
	assert.NoError(t, err)
	testCases(t, tc, admin)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAdministratorRepository_GetById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorById)).WithArgs(id).WillReturnError(ErrDatabaseAdministrator)

	admin, err := repo.GetById(context.Background(), id)

	assert.Nil(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetById_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	id := uuid.New()
	invalidEmail := "not-an-email"

	mockRows := sqlmock.NewRows([]string{
		"first_name", "last_name", "email", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		"John", "Doe", invalidEmail, "$2a$10$dummyhash", "male",
		time.Now(), nil, time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorById)).WithArgs(id).WillReturnRows(mockRows)

	admin, err := repo.GetById(context.Background(), id)

	assert.Nil(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	cases := Cases()
	tc := cases[3]

	mockRows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		tc.id, tc.firstName, tc.lastName, tc.password, tc.gender,
		tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorByEmail)).WithArgs(tc.email).WillReturnRows(mockRows)

	admin, err := repo.GetByEmail(context.Background(), tc.email)
	testCases(t, tc, admin)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAdministratorRepository_GetByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorByEmail)).WithArgs(email).WillReturnError(ErrDatabaseAdministrator)

	admin, err := repo.GetByEmail(context.Background(), email)

	assert.Nil(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_GetByEmail_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email := "valid@email.com"
	password := "soft"

	mockRows := sqlmock.NewRows([]string{
		"id", "first_name", "last_name", "password", "gender", "birth", "phone",
		"last_login_at", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		uuid.New(), "John", "Doe", password, "male",
		time.Now(), nil, time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetAdministratorByEmail)).WithArgs(email).WillReturnRows(mockRows)

	admin, err := repo.GetByEmail(context.Background(), email)

	assert.Nil(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_ExistById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	cases := Cases()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exists := tc.deletedAt == nil
			mockRow := sqlmock.NewRows([]string{"exists"}).AddRow(exists)

			mock.ExpectQuery(regexp.QuoteMeta(QueryExistAdministratorById)).WithArgs(tc.id).WillReturnRows(mockRow)

			result, err := repo.ExistById(context.Background(), tc.id)

			assert.NoError(t, err)
			assert.Equal(t, exists, result)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_ExistById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistAdministratorById)).WithArgs(id).WillReturnError(ErrDatabaseAdministrator)

	admin, err := repo.ExistById(context.Background(), id)

	assert.False(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_ExistByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	cases := Cases()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exists := tc.deletedAt == nil
			mockRow := sqlmock.NewRows([]string{"exists"}).AddRow(exists)

			mock.ExpectQuery(regexp.QuoteMeta(QueryExistAdministratorByEmail)).WithArgs(tc.email).WillReturnRows(mockRow)

			result, err := repo.ExistByEmail(context.Background(), tc.email)

			assert.NoError(t, err)
			assert.Equal(t, exists, result)
		})
	}

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_ExistByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistAdministratorByEmail)).WithArgs(email).WillReturnError(ErrDatabaseAdministrator)

	admin, err := repo.ExistByEmail(context.Background(), email)

	assert.False(t, admin)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		adm.Id(), adm.FirstName(), adm.LastName(), adm.Email().Value(),
		adm.Password().String(), adm.Gender().String(),
		adm.Birth().Value(), adm.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateAdministrator)).
		WithArgs(
			adm.Id(), adm.FirstName(), adm.LastName(),
			adm.Email().Value(), adm.Password().String(),
			string(adm.Gender()),
			adm.Birth().Value(),
			adm.Phone(),
		).WillReturnRows(rows)

	result, err := repo.Create(context.Background(), adm)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, adm.Id(), result.Id())
	assert.Equal(t, adm.FirstName(), result.FirstName())
	assert.Equal(t, adm.LastName(), result.LastName())
	assert.Equal(t, adm.Email(), result.Email())
	assert.Equal(t, adm.Password(), result.Password())
	assert.Equal(t, adm.Gender(), result.Gender())
	assert.Equal(t, adm.Birth(), result.Birth())
	assert.Equal(t, adm.Phone(), result.Phone())
	assert.NotNil(t, result.LastLoginAt())
	assert.NotNil(t, result.CreatedAt())
	assert.NotNil(t, result.UpdatedAt())
	assert.Nil(t, result.DeletedAt())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func TestAdministratorRepository_Create_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email, password, gender, birth := valueObjects(t, "test@example.com", "Str0ng!!1", "M", time.Now().AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateAdministrator)).
		WithArgs(
			adm.Id(), adm.FirstName(), adm.LastName(),
			adm.Email().Value(), adm.Password().String(),
			adm.Gender(),
			adm.Birth().Value(),
			adm.Phone(),
		).
		WillReturnError(ErrDatabaseAdministrator)

	result, err := repo.Create(context.Background(), adm)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrQueryAdministrator)
	assert.ErrorIs(t, err, ErrDatabaseAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Create_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email, password, gender, birth := valueObjects(t, "test@example.com", "Str0ng!!1", "M", time.Now().AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		adm.Id(), adm.FirstName(), adm.LastName(),
		adm.Email().Value(), adm.Password().String(),
		"X",
		adm.Birth().Value(),
		nil,
		time.Now(), time.Now(), time.Now(), nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateAdministrator)).
		WithArgs(
			adm.Id(), adm.FirstName(), adm.LastName(),
			adm.Email().Value(), adm.Password().String(),
			adm.Gender(),
			adm.Birth().Value(),
			adm.Phone(),
		).
		WillReturnRows(rows)

	result, err := repo.Create(context.Background(), adm)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)
	assert.Contains(t, err.Error(), "invalid")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		adm.Id(), adm.FirstName(), adm.LastName(), adm.Email().Value(),
		adm.Password().String(), adm.Gender().String(),
		adm.Birth().Value(), adm.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateAdministrator)).
		WithArgs(
			adm.FirstName(), adm.LastName(), adm.Email().Value(), adm.Password().String(),
			string(adm.Gender()), adm.Birth().Value(), adm.Phone(), adm.LastLoginAt(), adm.UpdatedAt(), adm.Id(),
		).WillReturnRows(rows)

	result, err := repo.Update(context.Background(), adm)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, adm.Id(), result.Id())
	assert.Equal(t, adm.FirstName(), result.FirstName())
	assert.Equal(t, adm.LastName(), result.LastName())
	assert.Equal(t, adm.Email().Value(), result.Email().Value())
	assert.Equal(t, adm.Password().String(), result.Password().String())
	assert.Equal(t, adm.Gender(), result.Gender())
	assert.Equal(t, adm.Birth(), result.Birth())
	assert.Equal(t, adm.Phone(), result.Phone())
	assert.NotNil(t, result.LastLoginAt())
	assert.NotNil(t, result.CreatedAt())
	assert.NotNil(t, result.UpdatedAt())
	assert.Nil(t, result.DeletedAt())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Update_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", time.Now().AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateAdministrator)).
		WithArgs(
			adm.FirstName(), adm.LastName(), adm.Email().Value(),
			adm.Password().String(), string(adm.Gender()), adm.Birth().Value(),
			adm.Phone(), adm.LastLoginAt(), adm.UpdatedAt(), adm.Id(),
		).
		WillReturnError(fmt.Errorf("database is down"))

	result, err := repo.Update(context.Background(), adm)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database is down")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Update_NewAdministratorError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)

	rows := sqlmock.NewRows(columns).AddRow(
		uuid.Nil, adm.FirstName(), adm.LastName(), "",
		adm.Password().String(), adm.Gender().String(),
		adm.Birth().Value(), adm.Phone(), now, now, now, nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateAdministrator)).
		WithArgs(
			adm.FirstName(), adm.LastName(), adm.Email().Value(),
			adm.Password().String(), string(adm.Gender()), adm.Birth().Value(),
			adm.Phone(), adm.LastLoginAt(), adm.UpdatedAt(), adm.Id(),
		).
		WillReturnRows(rows)

	result, err := repo.Update(context.Background(), adm)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "concatenating administrator values")
	assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdministratorRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)
	nColumns := []string{
		"first_name", "last_name", "email", "password", "gender",
		"birth", "phone", "last_login_at", "created_at", "updated_at", "deleted_at",
	}
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(nColumns).AddRow(
			adm.FirstName(), adm.LastName(), adm.Email().Value(),
			adm.Password().String(), adm.Gender().String(), adm.Birth().Value(),
			adm.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteAdministrator)).
			WithArgs(adm.Id()).
			WillReturnRows(rows)

		result, err := repo.Delete(context.Background(), adm.Id())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, adm.FirstName(), result.FirstName())
		assert.Equal(t, adm.LastName(), result.LastName())
		assert.Equal(t, adm.Email().Value(), result.Email().Value())
		assert.Equal(t, adm.Password().String(), result.Password().String())
		assert.Equal(t, adm.Gender(), result.Gender())
		assert.Equal(t, adm.Birth(), result.Birth())
		assert.Equal(t, adm.Phone(), result.Phone())
		assert.NotNil(t, result.DeletedAt())

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteAdministrator)).
			WithArgs(adm.Id()).
			WillReturnError(ErrDatabaseAdministrator)

		result, err := repo.Delete(context.Background(), adm.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryAdministrator)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("concatenation error", func(t *testing.T) {
		invalidEmail := ""
		rows := sqlmock.NewRows(nColumns).AddRow(
			adm.FirstName(), adm.LastName(), invalidEmail,
			adm.Password().String(), adm.Gender().String(), adm.Birth().Value(),
			adm.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteAdministrator)).
			WithArgs(adm.Id()).
			WillReturnRows(rows)

		result, err := repo.Delete(context.Background(), adm.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

}

func TestAdministratorRepository_Restore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)
	now := time.Now()
	email, password, gender, birth := valueObjects(t, "valid@email.com", "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa", "male", now.AddDate(-25, 0, 0))
	adm := administrators.NewAdministrator("John", "Doe", email, password, gender, birth, nil)
	nColumns := []string{
		"first_name", "last_name", "email", "password", "gender",
		"birth", "phone", "last_login_at", "created_at", "updated_at", "deleted_at",
	}
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(nColumns).AddRow(
			adm.FirstName(), adm.LastName(), adm.Email().Value(),
			adm.Password().String(), adm.Gender().String(), adm.Birth().Value(),
			adm.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryRestoreAdministrator)).
			WithArgs(adm.Id()).
			WillReturnRows(rows)

		result, err := repo.Restore(context.Background(), adm.Id())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, adm.FirstName(), result.FirstName())
		assert.Equal(t, adm.LastName(), result.LastName())
		assert.Equal(t, adm.Email().Value(), result.Email().Value())
		assert.Equal(t, adm.Password().String(), result.Password().String())
		assert.Equal(t, adm.Gender(), result.Gender())
		assert.Equal(t, adm.Birth(), result.Birth())
		assert.Equal(t, adm.Phone(), result.Phone())
		assert.NotNil(t, result.DeletedAt())

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryRestoreAdministrator)).
			WithArgs(adm.Id()).
			WillReturnError(ErrDatabaseAdministrator)

		result, err := repo.Restore(context.Background(), adm.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryAdministrator)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("concatenation error", func(t *testing.T) {
		invalidEmail := ""
		rows := sqlmock.NewRows(nColumns).AddRow(
			adm.FirstName(), adm.LastName(), invalidEmail,
			adm.Password().String(), adm.Gender().String(), adm.Birth().Value(),
			adm.Phone(), now, now, now, now,
		)

		mock.ExpectQuery(regexp.QuoteMeta(QueryRestoreAdministrator)).
			WithArgs(adm.Id()).
			WillReturnRows(rows)

		result, err := repo.Restore(context.Background(), adm.Id())
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrConcatenatingAdministrator)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

}

func TestAdministratorRepository_CountAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountAllAdministrators)).
			WillReturnRows(rows)

		count, err := repo.CountAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountAllAdministrators)).
			WillReturnError(ErrDatabaseAdministrator)

		count, err := repo.CountAll(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryAdministrator)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestAdministratorRepository_CountActive(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountActiveAdministrators)).WillReturnRows(rows)
		count, err := repo.CountActive(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountActiveAdministrators)).
			WillReturnError(ErrDatabaseAdministrator)

		count, err := repo.CountActive(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryAdministrator)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestAdministratorRepository_CountDeleted(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAdministratorRepository(db)

	t.Run("success", func(t *testing.T) {
		expectedCount := 5

		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountDeletedAdministrators)).WillReturnRows(rows)
		count, err := repo.CountDeleted(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(QueryCountDeletedAdministrators)).WillReturnError(ErrDatabaseAdministrator)
		count, err := repo.CountDeleted(context.Background())
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrQueryAdministrator)
		assert.Equal(t, -1, count)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func valueObjects(t *testing.T, email, password, gender string, birth time.Time) (vo.Email, vo.Password, vo.Gender, vo.BirthDate) {
	emailVo, err := vo.NewEmail(email)
	assert.NotEmpty(t, emailVo)
	assert.NoError(t, err)

	passwordVo, err := vo.NewHashedPassword(password)
	if err != nil {
		passwordVo, err = vo.NewPassword(password)
		assert.NotEmpty(t, passwordVo)
		assert.NoError(t, err)
	} else {
		assert.NotEmpty(t, passwordVo)
		assert.NoError(t, err)
	}

	genderVo, err := vo.ParseGender(gender)
	assert.NoError(t, err)

	birthVo, err := vo.NewBirthDate(birth)
	assert.NotEmpty(t, birthVo)
	assert.NoError(t, err)

	return emailVo, passwordVo, genderVo, birthVo
}

func testCases(t *testing.T, tc Case, admin *administrators.Administrator) {
	assert.NotNil(t, admin)
	assert.Equal(t, tc.id, admin.Id())
	assert.Equal(t, tc.firstName, admin.FirstName())
	assert.Equal(t, tc.lastName, admin.LastName())
	assert.Equal(t, tc.email, admin.Email().Value())
	assert.Equal(t, tc.password, admin.Password().String())
	assert.Equal(t, tc.gender, admin.Gender().String())
	assert.True(t, tc.birth.Equal(admin.Birth().Value()))
	if tc.phone != nil {
		phone, _ := vo.NewPhone(tc.phone)
		assert.NotNil(t, admin.Phone())
		assert.Equal(t, phone, admin.Phone())
	} else {
		assert.Nil(t, admin.Phone())
	}
	assert.True(t, tc.lastLoginAt.Equal(admin.LastLoginAt()))
	assert.True(t, tc.createdAt.Equal(admin.CreatedAt()))
	assert.True(t, tc.updatedAt.Equal(admin.UpdatedAt()))
	if tc.deletedAt != nil {
		assert.NotNil(t, admin.DeletedAt())
		assert.True(t, tc.deletedAt.Equal(*admin.DeletedAt()))
	} else {
		assert.Nil(t, admin.DeletedAt())
	}
}
