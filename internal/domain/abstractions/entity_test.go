package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEntity(t *testing.T) {
	id := uuid.New()

	entity := NewEntity(id)

	assert.NotNil(t, entity, "entity should not be nil")
	assert.Equal(t, id, entity.Id)
	assert.Empty(t, entity.DomainEvents(), "")
}
