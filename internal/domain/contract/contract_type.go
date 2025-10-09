package contracts

import "fmt"

type ContractType string

const (
	Monthly   ContractType = "M" // Monthly
	HalfMonth ContractType = "H" // Half-Month
)

func (t ContractType) String() string {
	switch t {
	case Monthly:
		return "monthly"
	case HalfMonth:
		return "half-month"
	default:
		return "unknown"
	}
}

func ParseContractType(s string) (ContractType, error) {
	switch s {
	case "monthly", "M":
		return Monthly, nil
	case "half-month", "H":
		return HalfMonth, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrTypeContract, s)
	}
}
