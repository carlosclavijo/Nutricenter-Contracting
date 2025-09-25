package commands

import "github.com/google/uuid"

type DeleteDeliveryCommand struct {
	ContractId    uuid.UUID
	DeliveryDayId uuid.UUID
}
