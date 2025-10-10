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
	assert.NotNil(t, entity.Id)
	assert.NotNil(t, entity.DomainEvents())
	assert.NotEmpty(t, entity)

	assert.Equal(t, id, entity.Id)

	assert.Empty(t, entity.DomainEvents())
}

func TestEntity_AddDomainEvents(t *testing.T) {
	id := uuid.New()
	entity := NewEntity(id)

	event := NewDomainEvent()
	entity.AddDomainEvent(*event)

	assert.Equal(t, entity.DomainEvents()[0].Id(), event.Id())

	assert.Len(t, entity.DomainEvents(), 1)
	assert.WithinDuration(t, entity.DomainEvents()[0].OccurredOn(), event.OccurredOn(), 1e6)
}

func TestEntity_ClearDomainEvents(t *testing.T) {
	id := uuid.New()
	entity := NewEntity(id)

	entity.AddDomainEvent(*NewDomainEvent())
	entity.ClearDomainEvents()

	assert.Nil(t, entity.DomainEvents())
	assert.Empty(t, entity.DomainEvents())
}
