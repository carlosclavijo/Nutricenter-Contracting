package contracts

import "fmt"

type ContractStatus string

const (
	Created   ContractStatus = "CREATED"
	Active    ContractStatus = "ACTIVE"
	Completed ContractStatus = "COMPLETED"
)

func (s ContractStatus) String() string {
	return string(s)
}

func ParseContractStatus(s string) (ContractStatus, error) {
	switch s {
	case "created":
		return Created, nil
	case "active":
		return Active, nil
	case "completed":
		return Completed, nil
	default:
		return "", fmt.Errorf("invalid contract status: %s", s)
	}
}
