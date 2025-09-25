package commands

import "github.com/google/uuid"

type UpdateDeliveryDayCommand struct {
	ContractId    uuid.UUID
	DeliveryDayId uuid.UUID
	Street        string
	Number        int
}
