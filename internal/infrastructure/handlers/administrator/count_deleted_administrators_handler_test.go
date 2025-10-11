package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAdministratorHandler_HandleCountDeleted(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cases := []struct {
		name      string
		count     int
		repoError error
		wantErr   error
	}{
		{"some deleted admins", 20, nil, nil},
		{"no deleted admins", 0, nil, nil},
		{"repository error", 0, ErrDbConnectionAdministrator, ErrDbConnectionAdministrator},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil

			mockRepo.On("CountDeleted", mock.Anything).Return(tc.count, tc.repoError)

			query := queries.CountDeletedAdministratorsQuery{}
			resp, err := handler.HandleCountDeleted(ctx, query)

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
