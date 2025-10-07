package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregateRoot(t *testing.T) {
	id := uuid.New()
	aggregate := NewAggregateRoot(id)

	assert.NotNil(t, aggregate)
	assert.NotNil(t, aggregate.Id)
	assert.NotEmpty(t, aggregate)

	assert.Equal(t, id, aggregate.Id)
}
