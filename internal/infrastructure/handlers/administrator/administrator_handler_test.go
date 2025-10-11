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
	"time"
)

var ErrDbConnectionAdministrator error = errors.New("db connection failed")

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

func (m *MockFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) (*administrators.Administrator, error) {
	args := m.Called(firstName, lastName, email, password, gender, birth, phone)
	return nil, args.Error(1)
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

func (m *MockRepository) Create(ctx context.Context, administrator *administrators.Administrator) (*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) Update(ctx context.Context, administrator *administrators.Administrator) (*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) Restore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	return nil, nil
}

func (m *MockRepository) CountAll(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockRepository) CountActive(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockRepository) CountDeleted(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Get(0).(int), args.Error(1)
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
