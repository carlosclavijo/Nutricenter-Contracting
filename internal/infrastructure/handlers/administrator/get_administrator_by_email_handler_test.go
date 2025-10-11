package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestAdministratorHandler_HandleGetByEmail(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cases := []struct {
		name, firstName, lastName, email, password, gender string
		birth                                              time.Time
	}{
		{"Case 1", "John", "Doe", "john@doe.com", "Strong1!", "male", time.Now().AddDate(-20, 0, 0)},
		{"Case 2", "Jane", "Doe", "jane@doe.com", "Not$0ftR", "female", time.Now().AddDate(-19, 0, 0)},
		{"Case 3", "Juan", "Perez", "juan@perez.com", "Ju4nPer3Z)", "male", time.Now().AddDate(-25, 10, 0)},
	}

	var admins []*administrators.Administrator
	for _, tc := range cases {
		email, password, gender, birth := valueObjects(t, tc.email, tc.password, tc.gender, tc.birth)
		admin := administrators.NewAdministrator(tc.firstName, tc.lastName, email, password, gender, birth, nil)
		admins = append(admins, admin)
	}

	mockRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(admins[0], nil)

	cmd := queries.GetAdministratorByEmailQuery{
		Email: admins[0].Email().Value(),
	}

	resp, err := handler.HandleGetByEmail(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.IsType(t, &dto.AdministratorDTO{}, resp)

	assert.Equal(t, admins[0].Id().String(), resp.Id)
	assert.Equal(t, admins[0].FirstName(), resp.FirstName)
	assert.Equal(t, admins[0].LastName(), resp.LastName)
	assert.Equal(t, admins[0].Email().Value(), resp.Email)
	assert.Equal(t, admins[0].Gender().String(), resp.Gender)
	assert.Equal(t, admins[0].Birth().Value(), resp.Birth)
	assert.Nil(t, resp.Phone)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestHandleGetByEmail_Error(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	mockRepo.On("GetByEmail", mock.Anything, mock.Anything).Return((*administrators.Administrator)(nil), ErrDbConnectionAdministrator)

	cmd := queries.GetAdministratorByEmailQuery{
		Email: "valid@email.com",
	}

	resp, err := handler.HandleGetByEmail(ctx, cmd)

	assert.ErrorIs(t, err, ErrDbConnectionAdministrator)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}
