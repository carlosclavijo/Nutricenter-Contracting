package queries

import "github.com/google/uuid"

type GetDeliveryByIdQuery struct {
	Id uuid.UUID
}
