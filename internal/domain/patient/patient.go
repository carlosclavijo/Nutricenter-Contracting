package patient

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"time"
)

type Patient struct {
	*abstractions.AggregateRoot
	Name      string
	Phone     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
