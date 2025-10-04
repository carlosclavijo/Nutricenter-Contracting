package contracts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContractFactory_Valid(t *testing.T) {
	factory := NewContractFactory()

	assert.Empty(t, factory)
	assert.NotNil(t, factory)

	_, ok := factory.(ContractFactory)
	assert.True(t, ok)
}
