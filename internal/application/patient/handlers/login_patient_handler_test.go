package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestPatientHandler_HandleLogin(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	cmd := commands.LoginPatientCommand{
		Email:    "login@email.valid",
		Password: "strong1!P",
	}

	email, err := vo.NewEmail(cmd.Email)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	password, err := vo.NewPassword(cmd.Password)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	password, err = vo.NewHashedPassword(string(hashedPassword))
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	gender, err := vo.ParseGender("F")
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	birth, _ := vo.NewBirthDate(time.Now().AddDate(-25, 0, 0))

	patient := patients.NewPatient("Jane", "Doe", email, password, gender, birth, nil)

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(true, nil)
	mockRepo.On("GetByEmail", mock.Anything, cmd.Email).Return(patient, nil)
	mockRepo.On("Update", mock.Anything, patient).Return(patient, nil)

	resp, err := handler.HandleLogin(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, cmd.Email, resp.Email)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestPatientHandler_HandleLogin_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.LoginPatientCommand{
		Email:    "john@doe.com",
		Password: "Str0ng!1",
	}

	email, err := vo.NewEmail(cmd.Email)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	password, err := vo.NewPassword(cmd.Password)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	password, err = vo.NewHashedPassword(string(hashedPassword))
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	gender, err := vo.ParseGender("M")
	assert.NoError(t, err)

	birth, err := vo.NewBirthDate(time.Now().AddDate(-25, 0, 0))
	assert.NotEmpty(t, birth)
	assert.NoError(t, err)

	patient := patients.NewPatient("John", "Doe", email, password, gender, birth, nil)

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(true, nil)
	mockRepo.On("GetByEmail", mock.Anything, cmd.Email).Return(patient, nil)
	mockRepo.On("Update", mock.Anything, patient).Return(nil, ErrDbFailurePatient)

	resp, err := handler.HandleLogin(ctx, cmd)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrDbFailurePatient)
	assert.Nil(t, resp)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestPatientHandler_HandleLogin_EmailError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.LoginPatientCommand{
		Email:    "",
		Password: "Strong1!",
	}

	resp, err := handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrEmptyEmail)
	assert.Nil(t, resp)

	cmd.Email = "invalid@email"
	resp, err = handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrInvalidEmail)
	assert.Nil(t, resp)

	cmd.Email = "user@subdomainaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"
	resp, err = handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongEmail)
	assert.Nil(t, resp)
}

func TestPatientHandler_HandleLogin_ExistenceCheck(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name        string
		emailExists bool
		repoError   error
		wantErr     error
	}{
		{"fail when repository returns error", false, ErrDbFailurePatient, ErrDbFailurePatient},
		{"fail when patient not found", false, nil, patients.ErrNotFoundPatient},
		{"success when patient exists", true, nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewPatientHandler(mockRepo, mockFactory)

			cmd := commands.LoginPatientCommand{
				Email:    "jane@example.com",
				Password: "Strong1!",
			}

			mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(tc.emailExists, tc.repoError)

			if tc.emailExists && tc.repoError == nil {
				email, err := vo.NewEmail(cmd.Email)
				assert.NotEmpty(t, email)
				assert.NoError(t, err)

				gender, err := vo.ParseGender("F")
				assert.NoError(t, err)

				birth, err := vo.NewBirthDate(time.Now().AddDate(-25, 0, 0))
				assert.NotEmpty(t, birth)
				assert.NoError(t, err)

				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
				assert.NoError(t, err)

				password, err := vo.NewPassword(string(hashedPassword))
				assert.NotEmpty(t, password)
				assert.NoError(t, err)

				patient := patients.NewPatient(
					"Jane",
					"Doe",
					email,
					password,
					gender,
					birth,
					nil,
				)

				mockRepo.On("GetByEmail", mock.Anything, email.Value()).Return(patient, nil)
				mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*patients.Patient")).Return(patient, nil)
			}

			resp, err := handler.HandleLogin(ctx, cmd)

			if tc.wantErr == nil {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
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

func TestPatientHandler_HandleLogin_PasswordError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.LoginPatientCommand{
		Email:    "jane@doe.com",
		Password: "",
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(true, nil)

	resp, err := handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrEmptyPassword)
	assert.Nil(t, resp)

	cmd.Password = "Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!"
	resp, err = handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongPassword)
	assert.Nil(t, resp)

	cmd.Password = "short"
	resp, err = handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrShortPassword)
	assert.Nil(t, resp)

	cmd.Password = "Abc123SSS"
	resp, err = handler.HandleLogin(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrSoftPassword)
	assert.Nil(t, resp)
}

func TestPatientHandler_HandleLogin_GetEmailError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.LoginPatientCommand{
		Email:    "jane@doe.com",
		Password: "5trong!S",
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(true, nil)
	mockRepo.On("GetByEmail", mock.Anything, cmd.Email).Return(nil, patients.ErrNotFoundPatient)

	resp, err := handler.HandleLogin(ctx, cmd)

	assert.ErrorIs(t, err, patients.ErrNotFoundPatient)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestPatientHandler_HandleLogin_InvalidPassword(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	handler := NewPatientHandler(mockRepo, nil)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Correct!Pass123"), bcrypt.DefaultCost)

	assert.NoError(t, err)

	patient, err := patients.NewPatientFromDB(uuid.New(), "Jane", "Doe", "jane@doe.com", string(hashedPassword), "female", time.Now().AddDate(-25, 0, 0), nil, time.Now(), time.Now(), time.Now(), nil)

	assert.NotEmpty(t, patient)
	assert.NoError(t, err)

	cmd := commands.LoginPatientCommand{
		Email:    "jane@doe.com",
		Password: "Wrong!Pass123",
	}

	mockRepo.On("ExistByEmail", mock.Anything, cmd.Email).Return(true, nil)
	mockRepo.On("GetByEmail", mock.Anything, cmd.Email).Return(patient, nil)

	resp, err := handler.HandleLogin(ctx, cmd)

	assert.Error(t, err)
	assert.ErrorIs(t, err, patients.ErrInvalidCredentialsPatient)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
}
