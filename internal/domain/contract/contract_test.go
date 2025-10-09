package contracts

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
			contract, err := NewContractFromDb(
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
			assert.Equal(t, tc.startDate.Format(time.RFC3339), contract.StartDate().Format(time.RFC3339))
			assert.Equal(t, tc.endDate.Format(time.RFC3339), contract.EndDate().Format(time.RFC3339))
			assert.Equal(t, tc.costValue, contract.CostValue())
			assert.Equal(t, tc.createdAt.Format(time.RFC3339), contract.CreatedAt().Format(time.RFC3339))
			assert.Equal(t, tc.updatedAt.Format(time.RFC3339), contract.UpdatedAt().Format(time.RFC3339))
			if tc.deletedAt != nil {
				assert.NotNil(t, contract.DeletedAt())
				assert.Equal(t, tc.deletedAt.Format(time.RFC3339), contract.DeletedAt().Format(time.RFC3339))
			} else {
				assert.Nil(t, contract.DeletedAt())
			}

			assert.NoError(t, err)
			assert.Empty(t, contract.Deliveries())
		})
	}
}

func TestNewContractFromDB_Invalid(t *testing.T) {
	id := uuid.New()
	administratorId := uuid.New()
	patientId := uuid.New()
	ctype := "X"
	status := "X"
	created := time.Now()
	start := time.Now().AddDate(0, 0, 5)
	end := time.Now().AddDate(0, 0, 20)
	cost := 1000
	ds := []deliveries.Delivery{}
	createdAt := time.Now().AddDate(0, -6, 0)
	updatedAt := time.Now().AddDate(0, -3, 0)

	contract, err := NewContractFromDb(id, administratorId, patientId, ctype, status, created, start, end, cost, ds, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, ErrTypeContract)
	assert.Nil(t, contract)

	ctype = "monthly"
	contract, err = NewContractFromDb(id, administratorId, patientId, ctype, status, created, start, end, cost, ds, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, ErrStatusContract)
	assert.Nil(t, contract)

	status = "created"

	contract, err = NewContractFromDb(id, administratorId, patientId, ctype, status, created, start, end, cost, ds, createdAt, updatedAt, nil)
	assert.NotNil(t, contract)
	assert.NoError(t, err)

	err = contract.Completed()
	assert.ErrorIs(t, err, ErrChangeStatusContract)

	err = contract.Active()
	assert.NoError(t, err)

	err = contract.Active()
	assert.ErrorIs(t, err, ErrChangeStatusContract)

	err = contract.Completed()
	assert.NoError(t, err)
}
