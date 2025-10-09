package contracts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContractType(t *testing.T) {
	monthly := Monthly
	halfMonth := HalfMonth
	other, err := ParseContractType("X")

	assert.NotNil(t, err)

	assert.Equal(t, "monthly", monthly.String())
	assert.Equal(t, "half-month", halfMonth.String())
	assert.Equal(t, "unknown", other.String())
	assert.Equal(t, Monthly, monthly)
	assert.Equal(t, HalfMonth, halfMonth)
	assert.NotEqual(t, ContractType(""), other.String())

	assert.ErrorIs(t, err, ErrTypeContract)

	ct, err := ParseContractType("monthly")
	assert.NoError(t, err)
	assert.Equal(t, Monthly, ct)

	ct, err = ParseContractType("M")
	assert.NoError(t, err)
	assert.Equal(t, Monthly, ct)

	ct, err = ParseContractType("half-month")
	assert.NoError(t, err)
	assert.Equal(t, HalfMonth, ct)

	ct, err = ParseContractType("H")
	assert.NoError(t, err)
	assert.Equal(t, HalfMonth, ct)

	ct, err = ParseContractType("invalid")
	assert.Error(t, err)
	assert.Equal(t, ContractType(""), ct)
}
