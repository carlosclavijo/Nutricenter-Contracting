package dto

import (
	administrator "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	patient "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"time"
)

type ContractResponse struct {
	ContractDTO
	AdministratorDTO *administrator.AdministratorDTO
	PatientDTO       *patient.PatientDTO
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}
