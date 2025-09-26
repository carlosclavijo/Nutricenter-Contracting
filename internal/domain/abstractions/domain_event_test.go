package abstractions

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewDomainEvent(t *testing.T) {
	event := NewDomainEvent()

	if event == nil {
		t.Fatal("Expected a domain, got nil")
	}

	if event.Id() == uuid.Nil {
		t.Error("expected non-nil UUID")
	}

	now := time.Now()
	if event.OccurredOn().After(now) {
		t.Error("expected event to occur after now")
	}

	if event.OccurredOn().Before(now.Add(-time.Second)) {
		t.Error("expected event to occur before now")
	}
}
