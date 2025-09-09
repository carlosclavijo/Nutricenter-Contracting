package dto

import "github.com/google/uuid"

type AdministratorDTO struct {
	Id    uuid.UUID
	Name  string
	Phone string
}
