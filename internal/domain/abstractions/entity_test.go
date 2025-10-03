package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEntity(t *testing.T) {
	id := uuid.New()

	entity := NewEntity(id)

	assert.NotNil(t, entity)
	assert.NotEmpty(t, entity)

	assert.NotNil(t, entity.Id)
	assert.Equal(t, id, entity.Id)

	assert.NotNil(t, entity.DomainEvents())
	assert.Empty(t, entity.DomainEvents())
}

func TestAddDomainEvents(t *testing.T) {
	id := uuid.New()
	entity := NewEntity(id)
	event := NewDomainEvent()

	entity.AddDomainEvent(*event)

	assert.Len(t, entity.DomainEvents(), 1)
	assert.Equal(t, event.Id(), entity.DomainEvents()[0].Id())
	assert.WithinDuration(t, event.OccurredOn(), entity.DomainEvents()[0].OccurredOn(), 1e6)
}

func TestClearDomainEvents(t *testing.T) {
	id := uuid.New()
	entity := NewEntity(id)
	entity.AddDomainEvent(*NewDomainEvent())

	entity.ClearDomainEvents()

	assert.Nil(t, entity.DomainEvents())
	assert.Empty(t, entity.DomainEvents())
}
