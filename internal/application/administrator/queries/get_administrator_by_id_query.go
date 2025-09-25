package queries

import "github.com/google/uuid"

type GetAdministratorByIdQuery struct {
	Id uuid.UUID
}
