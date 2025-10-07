package contracts

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContract_Status(t *testing.T) {
	coords, err := valueobjects.NewCoordinates(40.7128, -74.0060)
	assert.NoError(t, err)

	contract1, err := NewContractFactory().Create(uuid.New(), uuid.New(), HalfMonth, time.Now().AddDate(0, 0, 3), 100, "Main St", 123, coords)
	assert.NoError(t, err)

	contract2, err := NewContractFactory().Create(uuid.New(), uuid.New(), HalfMonth, time.Now().AddDate(0, 0, 3), 100, "Main St", 123, coords)
	assert.NoError(t, err)

	err = contract1.Active()
	assert.NoError(t, err)
	assert.Equal(t, Active, contract1.contractStatus)

	err = contract1.Active()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "only Created contracts can convert to Active")

	err = contract1.Completed()
	assert.NoError(t, err)
	assert.Equal(t, Finished, contract1.contractStatus)

	err = contract1.Active()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "only Created contracts can convert to Active")

	err = contract1.Completed()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "only Active contracts can convert to Finished")

	err = contract2.Completed()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "only Active contracts can convert to Finished")
}

func TestNewContractFromDb(t *testing.T) {
	cases := []struct {
		name                             string
		id, adminId, patientId           uuid.UUID
		cType, cStatus                   string
		creationDate, startDate, endDate time.Time
		costValue                        int
		createdAt, updatedAt             time.Time
		deletedAt                        *time.Time
	}{
		{"Case 1",
			uuid.New(), uuid.New(), uuid.New(),
			"M", "C",
			time.Now().AddDate(0, -1, 0),
			time.Now().AddDate(0, 0, 3),
			time.Now().AddDate(0, 0, 3).AddDate(0, 0, 29),
			200,
			time.Now().AddDate(0, -1, 0),
			time.Now().AddDate(0, -1, 0),
			nil,
		},
		{
			"Case 2",
			uuid.New(), uuid.New(), uuid.New(),
			"H", "A",
			time.Now().AddDate(0, -2, 0),
			time.Now().AddDate(0, 0, 5),
			time.Now().AddDate(0, 0, 5).AddDate(0, 0, 14),
			150,
			time.Now().AddDate(0, -2, 0),
			time.Now().AddDate(0, -2, 0),
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			contract := NewContractFromDb(
				tc.id, tc.adminId, tc.patientId, tc.cType, tc.cStatus,
				tc.creationDate, tc.startDate, tc.endDate, tc.costValue,
				[]deliveries.Delivery{}, tc.createdAt, tc.updatedAt, tc.deletedAt,
			)

			assert.NotNil(t, contract)

			assert.Equal(t, tc.id, contract.Id())
			assert.Equal(t, tc.adminId, contract.AdministratorId())
			assert.Equal(t, tc.patientId, contract.PatientId())
			assert.Equal(t, tc.cType, string(contract.ContractType()))
			assert.Equal(t, tc.cStatus, string(contract.ContractStatus()))
			assert.Equal(t, tc.creationDate.Format(time.RFC3339), contract.CreationDate().Format(time.RFC3339))
			assert.Equal(t, tc.startDate.Format("2006-01-02"), contract.StartDate().Format("2006-01-02"))
			assert.Equal(t, tc.endDate.Format("2006-01-02"), contract.EndDate().Format("2006-01-02"))
			assert.Equal(t, tc.costValue, contract.CostValue())
			assert.Equal(t, tc.createdAt.Format(time.RFC3339), contract.CreatedAt().Format(time.RFC3339))
			assert.Equal(t, tc.updatedAt.Format(time.RFC3339), contract.UpdatedAt().Format(time.RFC3339))
			if tc.deletedAt != nil {
				assert.NotNil(t, contract.DeletedAt())
				assert.Equal(t, tc.deletedAt.Format(time.RFC3339), contract.DeletedAt().Format(time.RFC3339))
			} else {
				assert.Nil(t, contract.DeletedAt())
			}

			assert.Empty(t, contract.Deliveries())
		})
	}
}
