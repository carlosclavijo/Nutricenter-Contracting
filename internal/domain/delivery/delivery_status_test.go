package deliveries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeliveryStatus_String(t *testing.T) {
	p := Pending
	d := Delivered
	c := Cancelled
	unknown := DeliveryStatus("")

	assert.Equal(t, p.String(), "pending")
	assert.Equal(t, d.String(), "delivered")
	assert.Equal(t, c.String(), "cancelled")
	assert.Equal(t, unknown.String(), "unknown")
}

func TestDeliveryStatus(t *testing.T) {
	cases := []struct {
		name, input string
		expected    DeliveryStatus
		wantErr     bool
	}{
		{"Pending", "pending", Pending, false},
		{"Delivered", "delivered", Delivered, false},
		{"Cancelled", "cancelled", Cancelled, false},
		{"Invalid", "unknown", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseDeliveryStatus(tc.input)

			if tc.wantErr {
				assert.NotNil(t, err)
				expected := "invalid delivery-status: " + tc.input
				assert.ErrorContains(t, err, expected)
				assert.Equal(t, DeliveryStatus(""), got)
			} else {
				assert.Nil(t, err)
				assert.NotEqual(t, DeliveryStatus(""), got)
				assert.Exactly(t, tc.expected, got)
			}
		})
	}
}
