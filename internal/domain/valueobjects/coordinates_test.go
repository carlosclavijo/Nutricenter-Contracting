package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCoordinates(t *testing.T) {
	cases := []struct {
		name     string
		lat, lon float64
	}{
		{"AlmostEdgeSouthWest", -90, -180},
		{"AlmostEdgeNorthWest", 90, -180},
		{"AlmostEdgeSouthEast", -90, 180},
		{"AlmostEdgeNorthEast", 90, 180},
		{"EquatorAndPrimeMeridian", 0, 0},
		{"NorthPole", 90, 0},
		{"SouthPole", -90, 0},
		{"GreenwichEast", 51.48, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coords, err := NewCoordinates(tc.lat, tc.lon)

			assert.NotNil(t, coords)

			assert.Equal(t, tc.lat, coords.Latitude())
			assert.Equal(t, tc.lon, coords.Longitude())

			assert.Nil(t, err)
			assert.NoError(t, err)
		})
	}
}

func TestNewCoordinates_LatitudeError(t *testing.T) {
	cases := []struct {
		name     string
		lat, lon float64
	}{
		{"LatitudeOutsideEast", -90.01, 0},
		{"LatitudeOutsideWest", 90.01, 0},
		{"HugeOutside", 500, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coords, err := NewCoordinates(tc.lat, tc.lon)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrOutOfBoundariesLatitude)

			assert.Empty(t, coords)
		})
	}
}

func TestNewCoordinates_LongitudeError(t *testing.T) {
	cases := []struct {
		name     string
		lat, lon float64
	}{
		{"LatitudeOutsideEast", 0, -180.01},
		{"LatitudeOutsideWest", 0, 180.01},
		{"HugeOutside", 0, 1000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coords, err := NewCoordinates(tc.lat, tc.lon)

			assert.NotNil(t, err)

			assert.ErrorIs(t, err, ErrOutOfBoundariesLongitude)

			assert.Empty(t, coords)
		})
	}
}
