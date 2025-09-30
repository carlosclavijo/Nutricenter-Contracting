package contracts

import (
	"fmt"
)

type ContractType string

const (
	Monthly   ContractType = "MONTHLY"
	HalfMonth ContractType = "HALF-MONTH"
)

func (t ContractType) String() string {
	return string(t)
}

func ParseContractType(s string) (ContractType, error) {
	switch s {
	case "monthly":
		return Monthly, nil
	case "half-month":
		return HalfMonth, nil
	default:
		return "", fmt.Errorf("invalid contract type '%s'", s)
	}
}
