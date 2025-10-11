package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestHandleUpdate(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewPatientHandler(mockRepo, mockFactory)
	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
		Phone:     nil,
	}

	var phone *vo.Phone
	if cmd.Phone != nil {
		phone, _ = vo.NewPhone(cmd.Phone)
	}

	email, password, gender, birth := valueObjects(t, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth)
	patient := patients.NewPatient(cmd.FirstName, cmd.LastName, email, password, gender, birth, phone)
	patient.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)
	mockRepo.On("Update", mock.Anything, patient).Return(patient, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, cmd.Id.String(), resp.Id)
	assert.Equal(t, cmd.FirstName, resp.FirstName)
	assert.Equal(t, cmd.LastName, resp.LastName)
	assert.Equal(t, cmd.Email, resp.Email)
	assert.Equal(t, cmd.Gender, resp.Gender)
	assert.Equal(t, cmd.Birth.Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
	assert.Equal(t, cmd.Phone, resp.Phone)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestHandleUpdate_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	email, password, gender, birth := valueObjects(t, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth)
	patient := patients.NewPatient(cmd.FirstName, cmd.LastName, email, password, gender, birth, nil)
	patient.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)
	mockRepo.On("Update", mock.Anything, patient).Return(nil, ErrDbFailurePatient)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, ErrDbFailurePatient)
	assert.Nil(t, resp)

	mockFactory.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestHandleUpdate_IdError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.Nil,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, patients.ErrEmptyIdPatient)
	assert.Nil(t, resp)
}

func TestHandleUpdate_EmailError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "",
		Password:  "Strong1!",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrEmptyEmail)
	assert.Nil(t, resp)

	cmd.Email = "invalid@email"
	resp, err = handler.HandleUpdate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrInvalidEmail)
	assert.Nil(t, resp)

	cmd.Email = "user@subdomainaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"
	resp, err = handler.HandleUpdate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongEmail)
	assert.Nil(t, resp)
}

func TestHandleUpdate_ExistenceCheck(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name      string
		idExists  bool
		repoError error
		wantErr   error
	}{
		{"id exists", true, nil, nil},
		{"id does not exist", false, nil, patients.ErrNotFoundPatient},
		{"repository returns error", false, ErrDbFailurePatient, ErrDbFailurePatient},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewPatientHandler(mockRepo, mockFactory)

			cmd := commands.UpdatePatientCommand{
				Id:        uuid.New(),
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "jane@example.com",
				Password:  "Strong1!",
				Gender:    "female",
				Birth:     time.Now().AddDate(-25, 0, 0),
				Phone:     nil,
			}

			mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(tc.idExists, tc.repoError)

			if tc.idExists && tc.repoError == nil {
				email, password, gender, birth := valueObjects(t, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth)
				patient := patients.NewPatient(cmd.FirstName, cmd.LastName, email, password, gender, birth, nil)
				patient.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

				mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*patients.Patient")).Return(patient, nil)
			}

			resp, err := handler.HandleUpdate(ctx, cmd)

			if tc.wantErr == nil {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, cmd.FirstName, resp.FirstName)
				assert.Equal(t, cmd.LastName, resp.LastName)
				assert.Equal(t, cmd.Email, resp.Email)
				assert.Equal(t, cmd.Gender, resp.Gender)
				assert.Equal(t, cmd.Birth.Format(time.RFC3339), resp.Birth.Format(time.RFC3339))
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

func TestHandleUpdate_PasswordError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "valid@email.com",
		Password:  "",
		Gender:    "female",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrEmptyPassword)
	assert.Nil(t, resp)

	cmd.Password = "Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!Abcdef1!"
	resp, err = handler.HandleUpdate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrLongPassword)
	assert.Nil(t, resp)

	cmd.Password = "short"
	resp, err = handler.HandleUpdate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrShortPassword)
	assert.Nil(t, resp)

	cmd.Password = "softpassword"
	resp, err = handler.HandleUpdate(ctx, cmd)
	assert.ErrorIs(t, err, vo.ErrSoftPassword)
	assert.Nil(t, resp)
}

func TestHandleUpdate_GenderError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "valid@email.com",
		Password:  "strong!1S",
		Gender:    "X",
		Birth:     time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrNotAGender)
	assert.Nil(t, resp)
}

func TestHandleUpdate_BirthError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "valid@email.com",
		Password:  "strong!1S",
		Gender:    "F",
		Birth:     time.Now().AddDate(5, 0, 0),
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrFutureDate)
	assert.Nil(t, resp)

	cmd.Birth = time.Now().AddDate(-10, 0, 0)
	resp, err = handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrUnderageDate)
	assert.Nil(t, resp)
}

func TestHandleUpdate_PhoneError(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	phone := "A787878"
	cmd := commands.UpdatePatientCommand{
		Id:        uuid.New(),
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "valid@email.com",
		Password:  "strong!1S",
		Gender:    "F",
		Birth:     time.Now().AddDate(-25, 0, 0),
		Phone:     &phone,
	}

	mockRepo.On("ExistById", mock.Anything, cmd.Id).Return(true, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrNotNumericPhoneNumber)
	assert.Nil(t, resp)

	phone = "787878"
	cmd.Phone = &phone
	resp, err = handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrShortPhoneNumber)
	assert.Nil(t, resp)

	phone = "787878787878"
	cmd.Phone = &phone
	resp, err = handler.HandleUpdate(ctx, cmd)

	assert.ErrorIs(t, err, vo.ErrLongPhoneNumber)
	assert.Nil(t, resp)
}
