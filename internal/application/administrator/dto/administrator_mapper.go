package dto

import administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"

func MapToAdministratorDTO(administrator administrators.Administrator) AdministratorDTO {
	return AdministratorDTO{
		Id:        administrator.Id().String(),
		FirstName: administrator.FirstName(),
		LastName:  administrator.LastName(),
		Email:     administrator.Email().Value(),
		Gender:    administrator.Gender().String(),
		Birth:     administrator.Birth().Value(),
		Phone:     administrator.Phone().String(),
	}
}
