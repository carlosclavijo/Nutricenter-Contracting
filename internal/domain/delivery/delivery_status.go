package deliveries

import "fmt"

type DeliveryStatus string

const (
	Pending   DeliveryStatus = "P" // Pending
	Delivered DeliveryStatus = "D" // Delivered
	Cancelled DeliveryStatus = "C" // Cancelled
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
	case "pending":
		return Pending, nil
	case "delivered":
		return Delivered, nil
	case "cancelled":
		return Cancelled, nil
	default:
		return "", fmt.Errorf("invalid delivery-status: %s", s)
	}
}
