package dto

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
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

func MapToAdministratorDTO(administrator *administrators.Administrator) *AdministratorDTO {
	if administrator == nil {
		return nil
	}

	return &AdministratorDTO{
		Id:        administrator.Id().String(),
		FirstName: administrator.FirstName(),
		LastName:  administrator.LastName(),
		Email:     administrator.Email().Value(),
		Gender:    administrator.Gender().String(),
		Birth:     administrator.Birth().Value(),
		Phone:     administrator.Phone().String(),
	}
}
