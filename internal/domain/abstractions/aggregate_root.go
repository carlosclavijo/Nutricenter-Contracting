package abstractions

import "github.com/google/uuid"

type AggregateRoot struct {
	*Entity
}

func NewAggregateRoot(id uuid.UUID) *AggregateRoot {
	return &AggregateRoot{NewEntity(id)}
}
