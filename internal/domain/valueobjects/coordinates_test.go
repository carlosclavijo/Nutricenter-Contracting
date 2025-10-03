package valueobjects

import (
	"fmt"
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
			assert.Equal(t, coords.Latitude(), tc.lat)
			assert.Equal(t, coords.Longitude(), tc.lon)

			assert.Nil(t, err)
			assert.NoError(t, err)
		})
	}
}

func TestNewCoordinates_InValid_Latitude(t *testing.T) {
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

			assert.Empty(t, coords)

			assert.NotNil(t, err)

			expected := fmt.Sprintf("latitude '%.2f' must been between -90 and 90", tc.lat)
			assert.ErrorContains(t, err, expected)
		})
	}
}

func TestNewCoordinates_InValid_Longitude(t *testing.T) {
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

			assert.Empty(t, coords)

			assert.NotNil(t, err)

			expected := fmt.Sprintf("longitude '%.2f' must been between -180 and 180", tc.lon)
			assert.ErrorContains(t, err, expected)
		})
	}
}
