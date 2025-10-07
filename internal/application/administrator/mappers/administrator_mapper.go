package mappers

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"time"
)

func MapToAdministratorDTO(administrator *administrators.Administrator) *dto.AdministratorDTO {
	var phone *string
	if administrator.Phone() != nil {
		p := administrator.Phone().String() // get string value
		phone = p
	}

	return &dto.AdministratorDTO{
		Id:        administrator.Id().String(),
		FirstName: administrator.FirstName(),
		LastName:  administrator.LastName(),
		Email:     administrator.Email().Value(),
		Gender:    administrator.Gender().String(),
		Birth:     administrator.Birth().Value(),
		Phone:     phone,
	}
}

func MapToAdministratorResponse(administrator *dto.AdministratorDTO, last, created, updated time.Time, deleted *time.Time) *dto.AdministratorResponse {
	return &dto.AdministratorResponse{
		AdministratorDTO: *administrator,
		LastLoginAt:      last,
		CreatedAt:        created,
		UpdatedAt:        updated,
		DeletedAt:        deleted,
	}
}
