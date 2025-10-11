package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestAdministratorHandler_HandleDelete(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	id := uuid.New()
	firstName := "John"
	lastName := "Doe"

	email, password, gender, birth := valueObjects(t, "valid@email.com", "sTrong!1s", "M", time.Now().AddDate(-20, 0, 0))

	phoneStr := "78787878"
	phone, err := vo.NewPhone(&phoneStr)
	assert.NotEmpty(t, phone)
	assert.NoError(t, err)

	admin := administrators.NewAdministrator(firstName, lastName, email, password, gender, birth, phone)

	mockRepo.On("ExistById", mock.Anything, id).Return(true, nil)
	mockRepo.On("Delete", mock.Anything, id).Return(admin, nil)

	resp, err := handler.HandleDelete(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Id)
	assert.Equal(t, firstName, resp.FirstName)
	assert.Equal(t, lastName, resp.LastName)
	assert.Equal(t, email.Value(), resp.Email)
	assert.Equal(t, gender.String(), resp.Gender)
	assert.Equal(t, birth.Value().Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
	assert.Equal(t, phone.String(), resp.Phone)
}

func TestAdministratorHandler_HandleDelete_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	id := uuid.New()

	mockRepo.On("ExistById", mock.Anything, id).Return(true, nil)
	mockRepo.On("Delete", mock.Anything, id).Return(nil, ErrDbFailureAdministrator)

	resp, err := handler.HandleDelete(ctx, id)

	assert.ErrorIs(t, err, ErrDbFailureAdministrator)
	assert.Nil(t, resp)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestAdministratorHandler_HandleDelete_IdError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	id := uuid.Nil
	resp, err := handler.HandleDelete(ctx, id)

	assert.ErrorIs(t, err, administrators.ErrEmptyIdAdministrator)
	assert.Nil(t, resp)
}

func TestAdministratorHandler_HandleDelete_ExistenceCheck(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name      string
		idExists  bool
		repoError error
		wantErr   error
	}{
		{"id exists", true, nil, nil},
		{"id does not exist", false, nil, administrators.ErrNotFoundAdministrator},
		{"repository returns error", false, ErrDbFailureAdministrator, ErrDbFailureAdministrator},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewAdministratorHandler(mockRepo, mockFactory)

			id := uuid.New()
			firstName := "John"
			lastName := "Doe"
			email := "valid@email.com"
			password := "5tronG.!!1"
			gender := "male"
			birth := time.Now().AddDate(-20, 0, 0)

			mockRepo.On("ExistById", mock.Anything, id).Return(tc.idExists, tc.repoError)

			if tc.idExists && tc.repoError == nil {
				nEmail, nPassword, nGender, nBirth := valueObjects(t, email, password, gender, birth)
				admin := administrators.NewAdministrator(firstName, lastName, nEmail, nPassword, nGender, nBirth, nil)
				admin.AggregateRoot = abstractions.NewAggregateRoot(id)

				mockRepo.On("Delete", mock.Anything, id).Return(admin, nil)
			}

			resp, err := handler.HandleDelete(ctx, id)

			if tc.wantErr == nil {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, firstName, resp.FirstName)
				assert.Equal(t, lastName, resp.LastName)
				assert.Equal(t, email, resp.Email)
				assert.Equal(t, gender, resp.Gender)
				assert.Equal(t, birth.Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, resp)
			}

			mockRepo.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}
