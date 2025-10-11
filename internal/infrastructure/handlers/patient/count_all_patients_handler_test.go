package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPatientHandler_HandleCountAll(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewPatientHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cases := []struct {
		name      string
		count     int
		repoError error
		wantErr   error
	}{
		{"some patients", 5, nil, nil},
		{"no patients", 0, nil, nil},
		{"repository error", 0, ErrDbConnectionPatient, ErrDbConnectionPatient},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil

			mockRepo.On("CountAll", mock.Anything).Return(tc.count, tc.repoError)

			query := queries.CountAllPatientsQuery{}
			resp, err := handler.HandleCountAll(ctx, query)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Equal(t, 0, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.count, resp)
			}

			mockRepo.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}
