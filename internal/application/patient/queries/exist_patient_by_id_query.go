package queries

import "github.com/google/uuid"

type ExistPatientByIdQuery struct {
	Id uuid.UUID
}
