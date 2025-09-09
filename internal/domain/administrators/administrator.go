package administrators

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/google/uuid"
)

type Administrator struct {
	*abstractions.AggregateRoot
	Name  string
	Phone string
}

func NewAdministrator(name, phone string) *Administrator {
	return &Administrator{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		Name:          name,
		Phone:         phone,
	}
}
