package contracts

import (
	"fmt"
)

type ContractType string

const (
	Monthly   ContractType = "monthly"
	HalfMonth ContractType = "half-month"
)

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
