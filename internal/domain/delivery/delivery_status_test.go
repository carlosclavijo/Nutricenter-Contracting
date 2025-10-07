package deliveries

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeliveryStatus(t *testing.T) {
	pending := Pending
	delivered := Delivered
	cancelled := Cancelled
	unknown, err := ParseDeliveryStatus("X")

	assert.Equal(t, "pending", pending.String())
	assert.Equal(t, "delivered", delivered.String())
	assert.Equal(t, "cancelled", cancelled.String())
	assert.Equal(t, "unknown", unknown.String())
	assert.Equal(t, Pending, pending)
	assert.Equal(t, Delivered, delivered)
	assert.Equal(t, Cancelled, cancelled)
	assert.NotEqual(t, DeliveryStatus(""), unknown.String())

	expected := fmt.Sprintf("input '%s' is not a delivery status", "X")
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expected)

	ds, err := ParseDeliveryStatus("pending")
	assert.NoError(t, err)
	assert.Equal(t, Pending, ds)

	ds, err = ParseDeliveryStatus("P")
	assert.NoError(t, err)
	assert.Equal(t, Pending, ds)

	ds, err = ParseDeliveryStatus("delivered")
	assert.NoError(t, err)
	assert.Equal(t, Delivered, ds)

	ds, err = ParseDeliveryStatus("D")
	assert.NoError(t, err)
	assert.Equal(t, Delivered, ds)

	ds, err = ParseDeliveryStatus("cancelled")
	assert.NoError(t, err)
	assert.Equal(t, Cancelled, ds)

	ds, err = ParseDeliveryStatus("C")
	assert.NoError(t, err)
	assert.Equal(t, Cancelled, ds)

	ds, err = ParseDeliveryStatus("invalid")
	assert.Equal(t, DeliveryStatus(""), ds)
	assert.Error(t, err)
}
