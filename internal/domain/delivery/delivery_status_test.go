package deliveries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeliveryStatus(t *testing.T) {
	pending := Pending
	delivered := Delivered
	cancelled := Cancelled
	unknown, err := ParseDeliveryStatus("X")

	assert.NotNil(t, err)

	assert.Equal(t, "pending", pending.String())
	assert.Equal(t, "delivered", delivered.String())
	assert.Equal(t, "cancelled", cancelled.String())
	assert.Equal(t, "unknown", unknown.String())
	assert.Equal(t, Pending, pending)
	assert.Equal(t, Delivered, delivered)
	assert.Equal(t, Cancelled, cancelled)
	assert.NotEqual(t, DeliveryStatus(""), unknown.String())

	ds, err := ParseDeliveryStatus("pending")
	assert.Equal(t, Pending, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("P")
	assert.Equal(t, Pending, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("delivered")
	assert.Equal(t, Delivered, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("D")
	assert.Equal(t, Delivered, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("cancelled")
	assert.Equal(t, Cancelled, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("C")
	assert.Equal(t, Cancelled, ds)
	assert.NoError(t, err)

	ds, err = ParseDeliveryStatus("invalid")
	assert.Equal(t, DeliveryStatus(""), ds)
	assert.ErrorIs(t, err, ErrNotADeliveryStatus)
}
