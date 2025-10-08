package deliveries

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestNewDeliveryFromDB(t *testing.T) {
	deletedAts := []time.Time{time.Now(), time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, -2, 0), time.Now().AddDate(0, -3, 0), time.Now().AddDate(0, -4, 0)}
	now := time.Now()
	cases := []struct {
		name, street, status       string
		number                     int
		lat, lon                   float64
		date, createdAt, updatedAt time.Time
		deletedAt                  *time.Time
	}{
		{"Case 1", "Sesame Street", "P", 30, -17.7863, -63.1812, now.AddDate(-1, -1, -7), now.AddDate(0, -1, -3), now.AddDate(-1, -6, -18), nil},
		{"Case 2", "Elm Street", "pending", 77, 48.8583701, 2.2944813, now.AddDate(-1, -2, -16), now.AddDate(0, -2, -15), now.AddDate(-1, -7, -3), &deletedAts[0]},
		{"Case 3", "Baker Street", "D", 912, 40.4319077, 116.5703749, now.AddDate(-1, -3, -11), now.AddDate(0, -3, -28), now.AddDate(-1, -8, -20), nil},
		{"Case 4", "Wall Street", "delivered", 10, 40.429657, 116.568115, now.AddDate(-1, -4, -25), now.AddDate(0, -4, -6), now.AddDate(-1, -9, -13), &deletedAts[3]},
		{"Case 5", "Diagon Alley", "C", 100, -17.7839, -63.1820, now.AddDate(-1, -5, -9), now.AddDate(0, -4, -6), now.AddDate(-1, -10, -27), nil},
		{"Case 6", "Abbey Road", "cancelled", 500, -18.1787, -63.7700, now.AddDate(-1, -6, -18), now.AddDate(0, -6, -12), now.AddDate(-1, -1, -7), &deletedAts[4]},
		{"Case 7", "Evergreen Terrace", "P", 93, -17.7818, -63.1800, now.AddDate(-1, -7, -3), now.AddDate(0, -7, -8), now.AddDate(-1, -2, -16), &deletedAts[2]},
		{"Case 8", "Wisteria Lane", "D", 1032, -17.8040, -63.1925, now.AddDate(-1, -8, -20), now.AddDate(0, -8, -23), now.AddDate(-1, -3, -11), nil},
		{"Case 9", "Electric Factory", "C", 44, -22.9519, -43.2105, now.AddDate(-1, -9, -13), now.AddDate(0, -9, -5), now.AddDate(-1, -4, -25), &deletedAts[1]},
		{"Case 10", "Penny Lane", "pending", 199, -33.8570, 151.2152, now.AddDate(-1, -10, -27), now.AddDate(0, -10, -14), now.AddDate(-1, -5, -9), nil},
	}

	for i, tc := range cases {
		c := 0
		var newStatus DeliveryStatus
		t.Run(tc.name, func(t *testing.T) {
			coords, err := valueobjects.NewCoordinates(tc.lat, tc.lon)

			assert.NoError(t, err)

			d := NewDelivery(uuid.New(), tc.date, tc.street, tc.number, coords)

			assert.NotNil(t, d.Id())
			assert.NotNil(t, d.ContractId())
			assert.NotEmpty(t, d.Coordinates())

			coords, err = valueobjects.NewCoordinates(tc.lat, tc.lon)

			assert.NoError(t, err)

			assert.Equal(t, tc.date.Format(time.RFC3339), d.Date().Format(time.RFC3339))
			assert.Equal(t, tc.street, d.Street())
			assert.Equal(t, tc.number, d.Number())
			assert.Equal(t, tc.lat, coords.Latitude())
			assert.Equal(t, tc.lon, coords.Longitude())
			assert.Equal(t, Pending, d.status)

			assert.Empty(t, d.CreatedAt())
			assert.Empty(t, d.UpdatedAt())
			assert.Empty(t, d.DeletedAt())

			d, err = NewDeliveryFromDB(uuid.New(), uuid.New(), tc.date, tc.street, tc.number, tc.lat, tc.lon, tc.status, tc.createdAt, tc.updatedAt, tc.deletedAt)

			assert.NotEmpty(t, d.Coordinates())

			assert.NotNil(t, d)
			assert.NotNil(t, d.Id())
			assert.NotNil(t, d.ContractId())

			status, err := ParseDeliveryStatus(tc.status)
			assert.NoError(t, err)

			coords, err = valueobjects.NewCoordinates(tc.lat, tc.lon)
			assert.NoError(t, err)

			assert.Equal(t, tc.date.Format(time.RFC3339), d.Date().Format(time.RFC3339))
			assert.Equal(t, tc.street, d.Street())
			assert.Equal(t, tc.number, d.Number())
			assert.Equal(t, tc.lat, coords.Latitude())
			assert.Equal(t, tc.lon, coords.Longitude())
			assert.Equal(t, status.String(), d.status.String())
			assert.Equal(t, tc.createdAt.Format(time.RFC3339), d.CreatedAt().Format(time.RFC3339))
			assert.Equal(t, tc.updatedAt.Format(time.RFC3339), d.UpdatedAt().Format(time.RFC3339))

			if tc.deletedAt != nil {
				assert.NotNil(t, d.DeletedAt())
				assert.Equal(t, tc.deletedAt.Format(time.RFC3339), d.DeletedAt().Format(time.RFC3339))
			} else {
				assert.Nil(t, d.DeletedAt())
			}

			if d.status == Pending {
				if c%2 == 0 {
					newStatus = Delivered
				} else {
					newStatus = Cancelled
				}

				street := "New Street" + strconv.Itoa(i)
				number := rand.Intn(1000) + 1
				lat := rand.Float64()*179 - 89
				lon := rand.Float64()*359 - 179
				coords, err = valueobjects.NewCoordinates(lat, lon)
				assert.NoError(t, err)

				oldUpdate := d.UpdatedAt()
				time.Sleep(10 * time.Millisecond)
				err := d.Update(street, number, coords)

				assert.NoError(t, err)

				assert.Equal(t, street, d.Street())
				assert.Equal(t, number, d.Number())
				assert.Equal(t, coords, d.Coordinates())

				coords = d.Coordinates()
				assert.Equal(t, lat, coords.Latitude())
				assert.Equal(t, lon, coords.Longitude())
				assert.True(t, d.UpdatedAt().After(oldUpdate))

				err = d.ChangeStatus(newStatus)

				assert.NoError(t, err)

				assert.Equal(t, newStatus, d.Status())

			} else {
				err := d.Update("", 10, valueobjects.Coordinates{})

				assert.NotNil(t, err)
				assert.ErrorIs(t, err, ErrNotPendingDelivery)

				err = d.ChangeStatus(newStatus)

				assert.NotNil(t, err)
				assert.ErrorIs(t, err, ErrCannotChangeDeliveryStatus)
			}
			c++
		})
	}
}
func TestNewDeliveryFromDB_Invalid(t *testing.T) {
	id := uuid.New()
	contractId := uuid.New()
	date := time.Now().AddDate(0, 0, 5)
	street := "New Street"
	number := 30
	lat := 91.0
	lon := -181.9
	status := "X"
	createdAt := time.Now().AddDate(0, 6, 0)
	updatedAt := time.Now().AddDate(0, 3, 0)

	delivery, err := NewDeliveryFromDB(id, contractId, date, street, number, lat, lon, status, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrOutOfBoundariesLatitude)
	assert.Nil(t, delivery)

	lat = 42
	delivery, err = NewDeliveryFromDB(id, contractId, date, street, number, lat, lon, status, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, valueobjects.ErrOutOfBoundariesLongitude)
	assert.Nil(t, delivery)

	lon = -90.48
	delivery, err = NewDeliveryFromDB(id, contractId, date, street, number, lat, lon, status, createdAt, updatedAt, nil)
	assert.ErrorIs(t, err, ErrNotADeliveryStatus)
	assert.Nil(t, delivery)
}
