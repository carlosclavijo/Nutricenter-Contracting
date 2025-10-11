package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAdministratorHandler_HandleExistByEmail(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockFactory := new(MockFactory)
	handler := NewAdministratorHandler(mockRepo, mockFactory)

	assert.NotNil(t, handler)

	cases := []struct {
		name               string
		exists, wantResp   bool
		repoError, wantErr error
	}{
		{"exists", true, true, nil, nil},
		{"does not exist", false, false, nil, nil},
		{"repository error", false, false, ErrDbConnectionAdministrator, ErrDbConnectionAdministrator},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil

			email := "valid@email.com"
			mockRepo.On("ExistByEmail", mock.Anything, email).Return(tc.exists, tc.repoError)

			cmd := queries.ExistAdministratorByEmailQuery{
				Email: email,
			}

			resp, err := handler.HandleExistByEmail(ctx, cmd)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.False(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResp, resp)
			}

			mockRepo.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}
