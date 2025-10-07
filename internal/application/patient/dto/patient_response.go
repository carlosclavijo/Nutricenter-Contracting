package dto

import "time"

type PatientResponse struct {
	PatientDTO
	LastLoginAt time.Time  `json:"last_login_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
