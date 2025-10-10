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

func TestAdministratorHandler_HandleCreate(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewAdministratorHandler(mockRepo, mockFactory)
	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
		Phone:     nil,
	}

	email, err := vo.NewEmail(cmd.Email)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	password, err := vo.NewPassword(cmd.Password)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	gender, err := vo.ParseGender(cmd.Gender)
	assert.NoError(t, err)

	birth, err := vo.NewBirthDate(cmd.Birth)
	assert.NotEmpty(t, birth)
	assert.NoError(t, err)

	admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, gender, birth, nil)

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, email, mock.Anything, gender, birth, (*vo.Phone)(nil)).Return(admin, nil)
	mockRepo.On("Create", mock.Anything, admin).Return(admin, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Id)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
	assert.Nil(t, resp.DeletedAt)

	assert.Equal(t, cmd.FirstName, resp.FirstName)
	assert.Equal(t, cmd.LastName, resp.LastName)
	assert.Equal(t, cmd.Email, resp.Email)
	assert.Equal(t, cmd.Gender, resp.Gender)
	assert.Equal(t, cmd.Birth.Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
	assert.Equal(t, cmd.Phone, resp.Phone)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestAdministratorHandler_HandleCreate_FactoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, vo.ErrInvalidEmail)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrInvalidEmail)
	assert.Nil(t, resp)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestAdministratorHandler_HandleCreate_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	email, err := vo.NewEmail(cmd.Email)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	password, err := vo.NewPassword(cmd.Password)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	gender, err := vo.ParseGender(cmd.Gender)
	assert.NoError(t, err)

	birth, err := vo.NewBirthDate(cmd.Birth)
	assert.NotEmpty(t, birth)
	assert.NoError(t, err)

	admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, gender, birth, nil)

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, email, mock.Anything, gender, birth, (*vo.Phone)(nil)).Return(admin, nil)
	mockRepo.On("Create", mock.Anything, admin).Return(nil, ErrDbFailureAdministrator)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, ErrDbFailureAdministrator)
	assert.Nil(t, resp)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestAdministratorHandler_HandleCreate_EmailError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "",
		Password:  "Strong1!",
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	resp, err := handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrEmptyEmail)
	assert.Nil(t, resp)

	cmd.Email = "invalid@email"
	resp, err = handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrInvalidEmail)
	assert.Nil(t, resp)

	cmd.Email = "user@subdomainaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"
	resp, err = handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongEmail)
	assert.Nil(t, resp)
}

func TestAdministratorHandler_HandleCreate_ExistenceCheck(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		emailExists bool
		repoError   error
		wantErr     error
	}{
		{"create when email does not exist", false, nil, nil},
		{"fail when email already exists", true, nil, administrators.ErrExistAdministrator},
		{"fail when repository returns error", false, ErrDbFailureAdministrator, ErrDbFailureAdministrator},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewAdministratorHandler(mockRepo, mockFactory)

			cmd := commands.CreateAdministratorCommand{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "jane@example.com",
				Password:  "Strong1!",
				Gender:    "female",
				Birth:     time.Now().AddDate(-25, 0, 0),
			}

			mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(tc.emailExists, tc.repoError)

			if !tc.emailExists && tc.repoError == nil {
				email, err := vo.NewEmail(cmd.Email)
				assert.NotEmpty(t, email)
				assert.NoError(t, err)

				password, err := vo.NewPassword(cmd.Password)
				assert.NotEmpty(t, password)
				assert.NoError(t, err)

				gender, err := vo.ParseGender(cmd.Gender)
				assert.NoError(t, err)

				birth, err := vo.NewBirthDate(cmd.Birth)
				assert.NotEmpty(t, birth)
				assert.NoError(t, err)

				admin := administrators.NewAdministrator(cmd.FirstName, cmd.LastName, email, password, gender, birth, nil)

				mockFactory.On("Create", cmd.FirstName, cmd.LastName, email, mock.Anything, gender, birth, (*vo.Phone)(nil)).Return(admin, nil)
				mockRepo.On("Create", mock.Anything, admin).Return(admin, nil)
			}

			resp, err := handler.HandleCreate(ctx, cmd)

			if tc.wantErr == nil {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			} else {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, resp)
			}

			mockRepo.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}

func TestAdministratorHandler_HandleCreate_PasswordError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "",
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)

	resp, err := handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrEmptyPassword)
	assert.Nil(t, resp)

	cmd.Password = "Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!"
	resp, err = handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongPassword)
	assert.Nil(t, resp)

	cmd.Password = "short"
	resp, err = handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrShortPassword)
	assert.Nil(t, resp)

	cmd.Password = "Abc123SSS"
	resp, err = handler.HandleCreate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrSoftPassword)
	assert.Nil(t, resp)
}

func TestAdministratorHandler_HandleCreate_GenderError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "Strong1!",
		Gender:    "X",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrNotAGender)
	assert.Nil(t, resp)
}

func TestAdministratorHandler_HandleCreate_BirthError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotEmpty(t, handler)

	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "Strong1!",
		Gender:    "male",
		Birth:     time.Now().AddDate(10, 0, 0),
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrFutureDate)
	assert.Empty(t, resp)

	cmd.Birth = time.Now().AddDate(-10, 0, 0)
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrUnderageDate)
	assert.Empty(t, resp)
}

func TestAdministratorHandler_HandleCreate_PhoneError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotEmpty(t, handler)

	phone := "78787878A"
	cmd := commands.CreateAdministratorCommand{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
		Password:  "Strong1!",
		Gender:    "male",
		Birth:     time.Now().AddDate(-25, 0, 0),
		Phone:     &phone,
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(false, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrNotNumericPhoneNumber)
	assert.Nil(t, resp)

	phone = "787878"
	cmd.Phone = &phone
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrShortPhoneNumber)
	assert.Nil(t, resp)

	phone = "78787878787878787878"
	cmd.Phone = &phone
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrLongPhoneNumber)
	assert.Nil(t, resp)
}
