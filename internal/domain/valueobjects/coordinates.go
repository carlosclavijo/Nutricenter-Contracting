package valueobjects

import (
	"errors"
	"fmt"
)

type Coordinates struct {
	lat float64
	lon float64
}

var (
	ErrOutOfBoundariesLatitude  = errors.New("latitude must be between -90 and 90")
	ErrOutOfBoundariesLongitude = errors.New("longitude must be between -180 and 180")
)

func NewCoordinates(latitude, longitude float64) (Coordinates, error) {
	if latitude < -90 || latitude > 90 {
		return Coordinates{}, fmt.Errorf("%w: got %.2f", ErrOutOfBoundariesLatitude, latitude)
	}
	if longitude < -180 || longitude > 180 {
		return Coordinates{}, fmt.Errorf("%w: got %.2f", ErrOutOfBoundariesLongitude, longitude)
	}
	return Coordinates{lat: latitude, lon: longitude}, nil
}

func (c Coordinates) Latitude() float64 {
	return c.lat
}

func (c Coordinates) Longitude() float64 {
	return c.lon
}
