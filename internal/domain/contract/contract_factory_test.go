package contracts

import (
	"fmt"
	deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewContractFactory(t *testing.T) {
	factory := NewContractFactory()

	assert.Empty(t, factory)
	assert.NotNil(t, factory)

	_, ok := factory.(ContractFactory)
	assert.True(t, ok)

	now := time.Now()

	cases := []struct {
		name, ctype, street string
		start               time.Time
		cost, number        int
		latitude, longitude float64
	}{
		{"Case 1", "monthly", "Sesame Street", now.AddDate(0, 0, 5), 5, 30, 40.7128, -74.0060},
		{"Case 2", "M", "Elm Street", now.AddDate(0, 0, 2), 100, 72, 35.6895, 139.6917},
		{"Case 3", "half-month", "Baker Street", now.AddDate(0, 0, 100), 1000, 129, 51.5074, -0.1278},
		{"Case 4", "H", "Wall Street", now.AddDate(0, 0, 50), 1500, 664, 48.8566, 2.3522},
		{"Case 5", "monthly", "Diagon Alley", now.AddDate(0, 0, 35), 7000, 700, -23.5505, -46.6333},
		{"Case 6", "M", "Abbey Road", now.AddDate(0, 0, 48), 2000, 100, 19.4326, -99.1332},
		{"Case 7", "half-month", "Evergreen Terrace", now.AddDate(0, 0, 29), 50, 90, 30.0444, -33.8688},
		{"Case 8", "H", "Wisteria Lane", now.AddDate(0, 0, 11), 1200, 358, 31.2357, 151.2093},
		{"Case 9", "monthly", "Electric Avenue", now.AddDate(0, 0, 2000), 1300, 902, 19.0760, 72.8777},
		{"Case 10", "half-month", "Penny Lane", now.AddDate(0, 0, 3), 80, 25, 41.0082, 28.9784},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			adminId, patientId := uuid.New(), uuid.New()

			ctype, err := ParseContractType(tc.ctype)

			assert.NoError(t, err)

			assert.NotEqual(t, "", ctype)

			coords, err := valueobjects.NewCoordinates(tc.latitude, tc.longitude)
			assert.NotEmpty(t, coords)
			assert.NoError(t, err)

			assert.Equal(t, coords.Latitude(), tc.latitude)
			assert.Equal(t, coords.Longitude(), tc.longitude)

			contract, err := factory.Create(adminId, patientId, ctype, tc.start, tc.cost, tc.street, tc.number, coords)
			assert.NotNil(t, contract)
			assert.NotNil(t, contract.Id())
			assert.NotNil(t, contract.EndDate())
			assert.NotEmpty(t, contract.EndDate())
			assert.NotEmpty(t, contract.Deliveries())
			assert.NoError(t, err)

			assert.Equal(t, adminId, contract.AdministratorId())
			assert.Equal(t, patientId, contract.PatientId())
			assert.Equal(t, ctype, contract.ContractType())
			assert.Contains(t, []string{"monthly", "half-month"}, contract.ContractType().String())
			assert.Equal(t, tc.start.Format(time.RFC3339), contract.StartDate().Format(time.RFC3339))
			assert.Equal(t, tc.cost, contract.CostValue())
			assert.Equal(t, Created, contract.ContractStatus())
			assert.WithinDuration(t, time.Now(), contract.CreationDate(), time.Second)

			days := 0
			if ctype == Monthly {
				days = 29
			} else if ctype == HalfMonth {
				days = 14
			}

			end := tc.start.AddDate(0, 0, days)
			assert.Equal(t, end, contract.EndDate())
			assert.Len(t, contract.Deliveries(), days+1)

			for i := 0; i < days; i++ {
				expectedDate := tc.start.AddDate(0, 0, i)
				actualDate := contract.Deliveries()[i].Date()
				d := contract.Deliveries()

				assert.NotNil(t, d[i])
				assert.NotNil(t, d[i].Id())
				assert.NotNil(t, d[i].ContractId())

				assert.Equal(t, contract.Id(), d[i].ContractId())
				assert.Equal(t, expectedDate.Format(time.RFC3339), actualDate.Format(time.RFC3339))
				assert.Equal(t, tc.street, d[i].Street())
				assert.Equal(t, tc.number, d[i].Number())

				dCoords := contract.Deliveries()[i].Coordinates()
				assert.NotEmpty(t, dCoords)
				assert.NotNil(t, d[i].CreatedAt())
				assert.NotNil(t, d[i].UpdatedAt())

				assert.Equal(t, dCoords, d[i].Coordinates())
				assert.Equal(t, tc.latitude, dCoords.Latitude())
				assert.Equal(t, tc.longitude, dCoords.Longitude())
				assert.Equal(t, deliveries.Pending, d[i].Status())
				assert.NotEqual(t, "", d[i].Status())

				assert.Empty(t, d[i].CreatedAt())
				assert.Empty(t, d[i].UpdatedAt())
				assert.Nil(t, d[i].DeletedAt())
			}

			assert.NotNil(t, contract.CreatedAt())
			assert.NotNil(t, contract.UpdatedAt())
			assert.Empty(t, contract.CreatedAt())
			assert.Empty(t, contract.UpdatedAt())

			assert.Nil(t, contract.DeletedAt())
		})
	}
}

func TestNewContractFactory_Errors(t *testing.T) {
	factory := NewContractFactory()

	administratorId := uuid.Nil
	patientId := uuid.New()
	contractType := Monthly
	start := time.Now().AddDate(0, 0, 5)
	cost := 100
	street := "Main Street"
	number := 123
	coordinates, err := valueobjects.NewCoordinates(40.7128, -74.0060)
	assert.NoError(t, err)

	contract, err := factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	assert.ErrorContains(t, err, "administratorId is not a valid UUID")
	assert.Nil(t, contract)

	administratorId = uuid.New()
	patientId = uuid.Nil

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	assert.ErrorContains(t, err, "patientId is not a valid UUID")
	assert.Nil(t, contract)

	patientId = uuid.New()
	contractType = ContractType("X")

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	expected := fmt.Sprintf("contractType '%s' is invalid", contractType)
	assert.ErrorContains(t, err, expected)
	assert.Nil(t, contract)

	contractType = HalfMonth
	start = time.Now().AddDate(0, 0, 1)

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	expected = fmt.Sprintf("startDate '%v' is not before two days after tomorrow", start)
	assert.ErrorContains(t, err, expected)
	assert.Nil(t, contract)

	start = time.Now().AddDate(0, 0, 5)
	cost = -10

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	expected = fmt.Sprintf("cost '%d' suppose to be a positive number", cost)
	assert.ErrorContains(t, err, expected)
	assert.Nil(t, contract)

	cost = 100
	street = ""

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	assert.ErrorContains(t, err, "street name is empty")
	assert.Nil(t, contract)

	street = "Main Street"
	number = 0

	contract, err = factory.Create(administratorId, patientId, contractType, start, cost, street, number, coordinates)
	expected = fmt.Sprintf("number '%d' needs to be a positive number", number)
	assert.ErrorContains(t, err, expected)
	assert.Nil(t, contract)
}
