package abstractions

import (
	"github.com/google/uuid"
)

type Entity struct {
	Id           uuid.UUID
	DomainEvents []DomainEvent
}

func NewEntity(id uuid.UUID) *Entity {
	return &Entity{
		Id:           id,
		DomainEvents: []DomainEvent{},
	}
}

func (e *Entity) AddDomainEvent(event DomainEvent) {
	e.DomainEvents = append(e.DomainEvents, event)
}

func (e *Entity) ClearDomainEvents() {
	e.DomainEvents = nil
}
