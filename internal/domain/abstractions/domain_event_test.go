package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDomainEvent(t *testing.T) {
	now := time.Now()

	event := NewDomainEvent()

	assert.NotNil(t, event, "NewDomainEvent returned nil event")
	assert.NotEqual(t, event.Id(), uuid.Nil, "NewDomainEvent returned nil event id")
	assert.WithinDuration(t, now, event.OccurredOn(), time.Second, "NewDomainEvent returned wrong time value")
}
