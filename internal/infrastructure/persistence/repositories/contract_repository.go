package repositories

import (
	"context"
	"database/sql"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"github.com/google/uuid"
)

type ContractRepository struct {
	DB *sql.DB
}

func (r *ContractRepository) GetList(ctx context.Context) (*[]dto.ContractDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContractRepository) GetById(ctx context.Context, id uuid.UUID) (*dto.ContractDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContractRepository) Create(ctx context.Context, c *contracts.Contract) (*contracts.Contract, error) {
	/*var (
		id, administratorId, patientId, dId uuid.UUID
		contractType                                          contracts.ContractType
		contractStatus                                        contracts.ContractStatus
		creation, start, end, createdAt, updatedAt, deletedAt time.Time
		cost                                                  int
		//deliveries                                            []deliveries.Delivery
	)

	query := `
		INSERT INTO contract(id, administrator_id, patient_id, contract_type, start_date, end_date, cost_value)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING *
	`

	err := r.DB.QueryRowContext(
		ctx, query, c.Id(), c.AdministratorId(), c.PatientId(), string(c.ContractType()), c.StartDate(), c.EndDate(), c.CostValue()).Scan(
		&id, &administratorId, &patientId, &contractType, &contractStatus, &creation, &start, &end, &cost, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:contract][Create] error executing SQL query '%s': '%v", query, err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	query = `
		INSERT INTO delivery(id, contract_id, date, street, number)
		VALUES($1, $2, $3, $4, $5)
		RETURNING *
	`

	for k, v := range c.Deliveries() {

	}

	contract := contracts.NewContractFromDb(id, administratorId, patientId, contractType, contractStatus, creation, start, &end, cost, createdAt, updatedAt, &deletedAt)
	log.Printf("[repository:contract][Create] successfully created contract %v", contract)
	return contract, nil*/
	//TODO implement me
	panic("implement me")
}

func (r *ContractRepository) Update(ctx context.Context, contract *contracts.Contract) (*contracts.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContractRepository) Delete(ctx context.Context, id uuid.UUID) (*contracts.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContractRepository) Restore(ctx context.Context, id uuid.UUID) (*contracts.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func NewContractRepository(db *sql.DB) contracts.ContractRepository {
	return &ContractRepository{DB: db}
}
