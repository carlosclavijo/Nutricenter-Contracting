package abstractions

import (
	"github.com/google/uuid"
	"time"
)

type DomainEvent struct {
	id         uuid.UUID
	occurredOn time.Time
}

func NewDomainEvent() *DomainEvent {
	return &DomainEvent{
		id:         uuid.New(),
		occurredOn: time.Now(),
	}
}

func (d *DomainEvent) Id() uuid.UUID {
	return d.id
}

func (d *DomainEvent) OccurredOn() time.Time {
	return d.occurredOn
}
