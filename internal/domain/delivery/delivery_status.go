package deliveries

import (
	"fmt"
)

type DeliveryStatus string

const (
	Pending   DeliveryStatus = "P"
	Delivered DeliveryStatus = "D"
	Cancelled DeliveryStatus = "C"
)

func (s DeliveryStatus) String() string {
	switch s {
	case Pending:
		return "pending"
	case Delivered:
		return "delivered"
	case Cancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

func ParseDeliveryStatus(s string) (DeliveryStatus, error) {
	switch s {
	case "pending", "P":
		return Pending, nil
	case "delivered", "D":
		return Delivered, nil
	case "cancelled", "C":
		return Cancelled, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrNotADeliveryStatus, s)
	}
}
