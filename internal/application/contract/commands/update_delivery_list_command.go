package commands

import (
	"github.com/google/uuid"
	"time"
)

type UpdateDeliveryDayListCommand struct {
	ContractId uuid.UUID
	FirstDate  time.Time
	LastDate   time.Time
	Street     string
	Number     int
	Latitude   float64
	Longitude  float64
}
