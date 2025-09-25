package queries

import "github.com/google/uuid"

type GetPatientByIdQuery struct {
	Id uuid.UUID
}
