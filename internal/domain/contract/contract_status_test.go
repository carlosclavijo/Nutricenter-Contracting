package contracts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContractStatus(t *testing.T) {
	created := Created
	active := Active
	finished := Finished
	other, err := ParseContractStatus("X")

	assert.NotNil(t, err)

	assert.Equal(t, created.String(), "created")
	assert.Equal(t, active.String(), "active")
	assert.Equal(t, finished.String(), "finished")
	assert.Equal(t, other.String(), "unknown")
	assert.Equal(t, Created, created)
	assert.Equal(t, Active, active)
	assert.Equal(t, Finished, finished)
	assert.NotEqual(t, ContractStatus(""), other.String())
	assert.ErrorIs(t, err, ErrStatusContract)

	cs, err := ParseContractStatus("created")
	assert.NoError(t, err)
	assert.Equal(t, Created, cs)

	cs, err = ParseContractStatus("C")
	assert.NoError(t, err)
	assert.Equal(t, Created, cs)

	cs, err = ParseContractStatus("active")
	assert.NoError(t, err)
	assert.Equal(t, Active, cs)

	cs, err = ParseContractStatus("A")
	assert.NoError(t, err)
	assert.Equal(t, Active, cs)

	cs, err = ParseContractStatus("finished")
	assert.NoError(t, err)
	assert.Equal(t, Finished, cs)

	cs, err = ParseContractStatus("F")
	assert.NoError(t, err)
	assert.Equal(t, Finished, cs)

	cs, err = ParseContractStatus("invalid")
	assert.Error(t, err)
	assert.Equal(t, ContractStatus(""), cs)
}
