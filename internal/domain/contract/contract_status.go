package contracts

import "fmt"

type ContractStatus string

const (
	Created    ContractStatus = "created"
	InProgress ContractStatus = "active"
	Completed  ContractStatus = "completed"
)

func ParseContractStatus(s string) (ContractStatus, error) {
	switch s {
	case "created":
		return Created, nil
	case "inprogress":
		return InProgress, nil
	case "completed":
		return Completed, nil
	default:
		return "", fmt.Errorf("invalid contract status: %s", s)
	}
}
