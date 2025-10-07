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
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	handler := NewAdministratorHandler(mockRepo, mockFactory)

	email, _ := vo.NewEmail(cmd.Email)
	passwordObj, _ := vo.NewPassword(cmd.Password)
	gender, _ := vo.ParseGender(cmd.Gender)
	birth, _ := vo.NewBirthDate(cmd.Birth)

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
	assert.Equal(t, cmd.FirstName, resp.FirstName)
	assert.Equal(t, cmd.LastName, resp.LastName)
	assert.Equal(t, cmd.Email, resp.Email)

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

	assert.Nil(t, resp)
	assert.Error(t, err)
}
