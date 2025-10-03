package abstractions

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDomainEvent(t *testing.T) {
	now := time.Now()

	event := NewDomainEvent()

	assert.NotNil(t, event)
	assert.NotEmpty(t, event)

	assert.NotNil(t, event.Id())
	
	assert.NotNil(t, event.OccurredOn())
	assert.WithinDuration(t, now, event.OccurredOn(), time.Second)
}
