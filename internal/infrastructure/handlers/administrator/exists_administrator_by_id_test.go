package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAdministratorHandler_HandleExistById(t *testing.T) {
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
			id := uuid.New()

			mockRepo.On("ExistById", mock.Anything, id).Return(tc.exists, tc.repoError)

			cmd := queries.ExistAdministratorByIdQuery{
				Id: id,
			}

			resp, err := handler.HandleExistById(ctx, cmd)

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
