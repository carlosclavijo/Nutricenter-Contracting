package commands

import (
	"github.com/google/uuid"
	"time"
)

type CreateContractCommand struct {
	AdministratorId uuid.UUID
	PatientId       uuid.UUID
	ContractType    string
	StartDate       time.Time
	Cost            int
	Street          string
	Number          int
	Latitude        float64
	Longitude       float64
}
