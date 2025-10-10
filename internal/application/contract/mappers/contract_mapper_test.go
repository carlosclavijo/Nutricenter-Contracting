package mappers

import (
	admMappers "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	ptnMappers "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapToContract_DTO_And_Response(t *testing.T) {
	administratorId := uuid.New()
	patientId := uuid.New()
	contractType := contracts.HalfMonth
	startDate := time.Now().AddDate(0, 0, 3)

	costValue := 1000
	street := "Elm Street"
	number := 30
	lat := -40.23
	lon := 72.83

	coordinates, err := valueobjects.NewCoordinates(lat, lon)
	assert.NoError(t, err)

	firstName := "John"
	lastName := "Doe"

	email, err := valueobjects.NewEmail("user@email.com")
	assert.NoError(t, err)

	password, err := valueobjects.NewPassword("Abc123!!")
	assert.NoError(t, err)

	gender := valueobjects.Male

	birth, err := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	assert.NoError(t, err)

	phoneStr := "78978978"
	phone, err := valueobjects.NewPhone(&phoneStr)
	assert.NoError(t, err)

	contract := contracts.NewContract(administratorId, patientId, contractType, startDate, costValue, street, number, coordinates)
	admin := administrators.NewAdministrator(firstName, lastName, email, password, gender, birth, phone)
	patient := patients.NewPatient(firstName, lastName, email, password, gender, birth, phone)

	contractDto := MapToContractDTO(contract)
	adminDto := admMappers.MapToAdministratorDTO(admin)
	patientDto := ptnMappers.MapToPatientDTO(patient)

	assert.NotNil(t, contractDto)
	assert.NotNil(t, adminDto)
	assert.NotNil(t, patientDto)

	assert.Equal(t, contract.Id().String(), contractDto.Id)
	assert.Equal(t, contract.AdministratorId().String(), contractDto.AdministratorId)
	assert.Equal(t, contract.PatientId().String(), contractDto.PatientId)
	assert.Equal(t, contract.ContractType().String(), contractDto.ContractType)
	assert.Equal(t, contract.ContractStatus().String(), contractDto.ContractStatus)
	assert.Equal(t, contract.CreationDate().Format(time.RFC3339), contractDto.CreationDate.Format(time.RFC3339))
	assert.Equal(t, contract.StartDate().Format(time.RFC3339), contractDto.StartDate.Format(time.RFC3339))
	assert.Equal(t, contract.EndDate().Format(time.RFC3339), contractDto.EndDate.Format(time.RFC3339))
	assert.Equal(t, contract.CostValue(), contractDto.CostValue)

	var deliveryDtos []*dto.DeliveryDTO
	for _, d := range contract.Deliveries() {
		dDto := MapToDeliveryDTO(&d)
		deliveryDtos = append(deliveryDtos, dDto)
	}

	assert.Exactly(t, deliveryDtos, contractDto.Deliveries)

	response := MapToContractResponse(contractDto, adminDto, patientDto, contract.CreatedAt(), contract.UpdatedAt(), contract.DeletedAt())

	assert.NotNil(t, response)

	assert.Exactly(t, *contractDto, response.ContractDTO)
	assert.Exactly(t, adminDto, response.AdministratorDTO)
	assert.Exactly(t, patientDto, response.PatientDTO)
	assert.Equal(t, contract.CreatedAt(), response.CreatedAt)
	assert.Equal(t, contract.UpdatedAt(), response.UpdatedAt)
	assert.Equal(t, contract.DeletedAt(), response.DeletedAt)
}
