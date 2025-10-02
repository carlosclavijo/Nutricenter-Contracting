package contracts

import "fmt"

type ContractStatus string

const (
	Created  ContractStatus = "C" // Created
	Active   ContractStatus = "A" // Active
	Finished ContractStatus = "F" // Finished
)

func (s ContractStatus) String() string {
	switch s {
	case Created:
		return "created"
	case Active:
		return "active"
	case Finished:
		return "finished"
	default:
		return "unknown"
	}
}

func ParseContractStatus(s string) (ContractStatus, error) {
	switch s {
	case "created", "C":
		return Created, nil
	case "active", "A":
		return Active, nil
	case "finished", "F":
		return Finished, nil
	default:
		return "", fmt.Errorf("invalid contract status: %s", s)
	}
}
