package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestPatientHandler_HandleGetAll(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cases := []struct {
		name, firstName, lastName, email, password, gender string
		birth                                              time.Time
	}{
		{"Case 1", "John", "Doe", "john@doe.com", "Strong1!", "male", time.Now().AddDate(-20, 0, 0)},
		{"Case 2", "Jane", "Doe", "jane@doe.com", "Not$0ftR", "female", time.Now().AddDate(-19, 0, 0)},
		{"Case 3", "Juan", "Perez", "juan@perez.com", "Ju4nPer3Z)", "male", time.Now().AddDate(-25, 10, 0)},
	}

	var ptnts []*patients.Patient

	for _, tc := range cases {
		email, password, gender, birth := valueObjects(t, tc.email, tc.password, tc.gender, tc.birth)
		patient := patients.NewPatient(tc.firstName, tc.lastName, email, password, gender, birth, nil)
		ptnts = append(ptnts, patient)
	}

	mockRepo.On("GetAll", mock.Anything).Return(ptnts, nil)

	cmd := queries.GetAllPatientsQuery{}

	resp, err := handler.HandleGetAll(ctx, cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.IsType(t, []*dto.PatientDTO{}, resp)
	assert.Len(t, resp, len(cases))

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := resp[i]
			assert.NotNil(t, got.Id)
			assert.Equal(t, tc.firstName, got.FirstName)
			assert.Equal(t, tc.lastName, got.LastName)
			assert.Equal(t, tc.email, got.Email)
			assert.Equal(t, tc.gender, got.Gender)
			assert.Equal(t, tc.birth, got.Birth)

			assert.Nil(t, got.Phone)
		})
	}

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestHandleGetAll_Error(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	mockRepo.On("GetAll", mock.Anything).Return([]*patients.Patient(nil), ErrDbConnectionPatient)

	cmd := queries.GetAllPatientsQuery{}

	resp, err := handler.HandleGetAll(ctx, cmd)

	assert.ErrorIs(t, err, ErrDbConnectionPatient)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}
