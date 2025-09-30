package valueobjects

import (
	"fmt"
)

type Coordinates struct {
	lat float64
	lon float64
}

func NewCoordinates(latitude, longitude float64) (*Coordinates, error) {
	if latitude < -90 || latitude > 90 {
		return nil, fmt.Errorf("latitude '%.2f' must be between -90 and 90", latitude)
	}
	if longitude < -180 || longitude > 180 {
		return nil, fmt.Errorf("longitude '%.2f' must be between -180 and 180", longitude)
	}
	return &Coordinates{lat: latitude, lon: longitude}, nil
}

func (c *Coordinates) Latitude() float64 {
	return c.lat
}

func (c *Coordinates) Longitude() float64 {
	return c.lon
}
