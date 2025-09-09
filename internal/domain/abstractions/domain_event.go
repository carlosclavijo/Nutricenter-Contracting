package abstractions

import (
	"github.com/google/uuid"
	"time"
)

type DomainEvent struct {
	Id        uuid.UUID
	OcurredOn time.Time
}

func NewDomainEvent() *DomainEvent {
	return &DomainEvent{
		Id:        uuid.New(),
		OcurredOn: time.Now(),
	}
}
