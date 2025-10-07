package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type MockFactory struct {
	mock.Mock
}

func (m *MockFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*administrators.Administrator, error) {
	args := m.Called(firstName, lastName, email, password, gender, birth, phone)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) GetAll(ctx context.Context) ([]*administrators.Administrator, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) GetList(ctx context.Context) ([]*administrators.Administrator, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) GetById(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) GetByEmail(ctx context.Context, email string) (*administrators.Administrator, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, admin *administrators.Administrator) (*administrators.Administrator, error) {
	args := m.Called(ctx, admin)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, admin *administrators.Administrator) (*administrators.Administrator, error) {
	args := m.Called(ctx, admin)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) Restore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*administrators.Administrator), args.Error(1)
}

func (m *MockRepository) CountAll(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) CountActive(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) CountDeleted(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}
