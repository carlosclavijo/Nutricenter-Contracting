package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var ErrDbFailureAdministrator error = errors.New("db failure")

type MockRepository struct {
	mock.Mock
}

type MockFactory struct {
	mock.Mock
}

func TestNewAdministratorHandler(t *testing.T) {
	r := new(MockRepository)
	f := new(MockFactory)
	h := NewAdministratorHandler(r, f)

	assert.NotEmpty(t, h)
}

func (m *MockFactory) Create(firstName string, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*administrators.Administrator, error) {
	args := m.Called(firstName, lastName, email, password, gender, birth, phone)

	var result *administrators.Administrator
	if v := args.Get(0); v != nil {
		result = v.(*administrators.Administrator)
	}

	return result, args.Error(1)
}

func (m *MockRepository) GetAll(ctx context.Context) ([]*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) GetList(ctx context.Context) ([]*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) GetById(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) GetByEmail(ctx context.Context, email string) (*administrators.Administrator, error) {
	args := m.Called(ctx, email)
	if v := args.Get(0); v != nil {
		return v.(*administrators.Administrator), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	args := m.Called(ctx, id)

	var exists bool
	if v := args.Get(0); v != nil {
		exists, _ = v.(bool)
	}

	return exists, args.Error(1)
}

func (m *MockRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	v := args.Get(0)
	if v == nil {
		return false, args.Error(1)
	}
	return v.(bool), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, admin *administrators.Administrator) (*administrators.Administrator, error) {
	args := m.Called(ctx, admin)

	var result *administrators.Administrator
	if v := args.Get(0); v != nil {
		result = v.(*administrators.Administrator)
	}

	return result, args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, admin *administrators.Administrator) (*administrators.Administrator, error) {
	args := m.Called(ctx, admin)

	var result *administrators.Administrator
	if v := args.Get(0); v != nil {
		result = v.(*administrators.Administrator)
	}

	return result, args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)

	var admin *administrators.Administrator
	if v := args.Get(0); v != nil {
		admin, _ = v.(*administrators.Administrator)
	}

	return admin, args.Error(1)
}

func (m *MockRepository) Restore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)

	var admin *administrators.Administrator
	if v := args.Get(0); v != nil {
		admin, _ = v.(*administrators.Administrator)
	}

	return admin, args.Error(1)
}

func (m *MockRepository) CountAll(ctx context.Context) (int, error) {
	return 0, nil
}

func (m *MockRepository) CountActive(ctx context.Context) (int, error) {
	return 0, nil
}

func (m *MockRepository) CountDeleted(ctx context.Context) (int, error) {
	return 0, nil
}
