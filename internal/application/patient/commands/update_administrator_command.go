package commands

import (
	"github.com/google/uuid"
	"time"
)

type UpdatePatientCommand struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    string
	Birth     *time.Time
	Phone     *string
}
