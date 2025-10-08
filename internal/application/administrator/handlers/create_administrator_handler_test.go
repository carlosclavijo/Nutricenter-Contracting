package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestHandleCreate(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	handler := NewAdministratorHandler(mockRepo, mockFactory)
	assert.NotEmpty(t, handler)

	email, err := vo.NewEmail(cmd.Email)
	assert.NoError(t, err)

	passwordObj, err := vo.NewPassword(cmd.Password)
	assert.NoError(t, err)

	gender, err := vo.ParseGender(cmd.Gender)
	assert.NoError(t, err)

	birth, err := vo.NewBirthDate(cmd.Birth)
	assert.NoError(t, err)

	adminObj := administrators.NewAdministrator(
		cmd.FirstName,
		cmd.LastName,
		email,
		passwordObj,
		gender,
		birth,
		nil,
	)

	mockFactory.
		On("Create", cmd.FirstName, cmd.LastName, email, mock.Anything, gender, birth, (*vo.Phone)(nil)).
		Return(adminObj, nil)

	mockRepo.
		On("Create", ctx, adminObj).
		Return(adminObj, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Id)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)

	assert.Equal(t, cmd.FirstName, resp.FirstName)
	assert.Equal(t, cmd.LastName, resp.LastName)
	assert.Equal(t, cmd.Email, resp.Email)
	assert.Equal(t, cmd.Gender, resp.Gender)
	assert.Equal(t, cmd.Birth.Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
	assert.Equal(t, cmd.Phone, resp.Phone)

	assert.Nil(t, resp.DeletedAt)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestHandleCreate_EmailError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "invalid-email",
		Password:  "Strong1!",
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	handler := NewAdministratorHandler(mockRepo, mockFactory)
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.NotNil(t, handler)

	assert.Error(t, err)

	assert.Nil(t, resp)
}

func TestHandleCreate_PasswordError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "weak",
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	handler := NewAdministratorHandler(mockRepo, mockFactory)
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.NotNil(t, handler)

	assert.Error(t, err)

	assert.Nil(t, resp)

}

func TestHandleCreate_GenderError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "Strong1!",
		Gender:    "X",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	handler := NewAdministratorHandler(mockRepo, mockFactory)
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.NotNil(t, handler)

	assert.Error(t, err)

	assert.Nil(t, resp)
}
