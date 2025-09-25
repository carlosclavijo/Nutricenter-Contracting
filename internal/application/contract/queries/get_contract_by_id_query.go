package queries

import "github.com/google/uuid"

type GetContractByIdQuery struct {
	Id uuid.UUID
}
