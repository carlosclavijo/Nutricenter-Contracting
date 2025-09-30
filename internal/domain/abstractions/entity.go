package abstractions

import (
	"github.com/google/uuid"
)

type Entity struct {
	Id           uuid.UUID
	domainEvents []DomainEvent
}

func NewEntity(id uuid.UUID) *Entity {
	return &Entity{
		Id:           id,
		domainEvents: []DomainEvent{},
	}
}

func (e *Entity) DomainEvents() []DomainEvent {
	return e.domainEvents
}

func (e *Entity) AddDomainEvent(event DomainEvent) {
	e.domainEvents = append(e.domainEvents, event)
}

func (e *Entity) ClearDomainEvents() {
	e.domainEvents = nil
}
