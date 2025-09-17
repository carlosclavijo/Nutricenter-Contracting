package dto

import (
	"time"
)

type AdministratorDTO struct {
	Id        string     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Gender    string     `json:"gender"`
	Birth     *time.Time `json:"birth,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
}
