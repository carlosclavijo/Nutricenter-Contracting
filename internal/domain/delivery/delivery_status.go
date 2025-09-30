package deliveries

import "fmt"

type DeliveryStatus string

const (
	Pending   DeliveryStatus = "PENDING"
	Delivered DeliveryStatus = "DELIVERED"
	Cancelled DeliveryStatus = "CANCELLED"
)

func (s DeliveryStatus) String() string {
	return string(s)
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
