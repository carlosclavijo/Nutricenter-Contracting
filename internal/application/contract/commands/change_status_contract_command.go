package commands

import "github.com/google/uuid"

type ChangeStatusContractCommand struct {
	Id     uuid.UUID
	Status string
}
