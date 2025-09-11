package administrator

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/google/uuid"
	"time"
)

type Administrator struct {
	*abstractions.AggregateRoot
	Name      string
	Phone     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewAdministrator(name, phone string) *Administrator {
	return &Administrator{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		Name:          name,
		Phone:         phone,
		CreatedAt:     nil,
		UpdatedAt:     nil,
	}
}
