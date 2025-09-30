package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

type ContractRepository struct {
	DB *sql.DB
}

func (r *ContractRepository) GetAll(ctx context.Context) (*[]dto.ContractDTO, error) {
	var cs []dto.ContractDTO

	query := `
		SELECT id, administrator_id, patient_id, type, status, creation, start, finalized, cost
		FROM contract
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:contract][GetAll] error executing SQL statement: %v", err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("[repository:contract][GetAll] error closing rows: %v]")
			return
		}

	}(rows)
	for rows.Next() {
		var c dto.ContractDTO
		err = rows.Scan(&c.Id, &c.AdministratorId, &c.PatientId, &c.ContractType, &c.ContractStatus, &c.CreationDate, &c.StartDate, &c.EndDate, &c.CostValue)
		if err != nil {
			log.Printf("[repository:contract][GetAll] error scanning rows: %v", err)
			return nil, fmt.Errorf("rows scan failed: %w", err)
		}
		cs = append(cs, c)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:contract][GetAll] error scanning rows: %v", err)
		return nil, fmt.Errorf("rows scan failed: %w", err)
	}

	log.Printf("[repository:contract][GetAll] successfully fetched %d")
	return &cs, nil
}

func (r *ContractRepository) GetById(ctx context.Context, id uuid.UUID) (*dto.ContractDTO, error) {
	var c dto.ContractDTO

	query := `
		SELECT id, administrator_id, patient_id, type, status, creation, start, finalized, cost
		FROM contract
		WHERE id = $1
	`

	err := r.DB.QueryRowContext(ctx, query, id).Scan(&c.Id, &c.AdministratorId, &c.PatientId, &c.ContractType, &c.ContractStatus, &c.CreationDate, &c.StartDate, &c.EndDate, &c.CostValue)
	if err != nil {
		log.Printf("[repository:contract][GetById] error scanning rows: %v", err)
		return nil, fmt.Errorf("rows scan failed: %w", err)
	}

	log.Printf("[repository:contract][GetById] successfully fetched")
	return &c, nil
}

func (r *ContractRepository) Create(ctx context.Context, c *contracts.Contract) (*contracts.Contract, error) {
	var (
		id, administratorId, patientId             uuid.UUID
		contractType                               contracts.ContractType
		contractStatus                             contracts.ContractStatus
		creation, start, end, createdAt, updatedAt time.Time
		deletedAt                                  *time.Time
		cost                                       int
	)

	query := `
		INSERT INTO contract(id, administrator_id, patient_id, type, start, finalized, cost)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, administrator_id, patient_id, type, status, creation, start, finalized, cost, created_at, updated_at, deleted_at
	`

	err := r.DB.QueryRowContext(
		ctx, query,
		c.Id(), c.AdministratorId(), c.PatientId(),
		string(c.ContractType()), c.StartDate(), c.EndDate(), c.CostValue(),
	).Scan(
		&id, &administratorId, &patientId, &contractType, &contractStatus,
		&creation, &start, &end, &cost, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:contract][Create] failed inserting contract: %v", err)
		return nil, fmt.Errorf("contract insert failed: %w", err)
	}

	d := c.Deliveries()
	if len(d) != 15 && len(d) != 30 {
		log.Printf("[repository:contract][Create] contract deliveries length: %v cannot be other than 15 or 30", len(d))
		return nil, fmt.Errorf("deliveries can only be 15 or 30 long")
	}

	var placeholders []string
	var args []interface{}
	for i, d := range d {
		base := i * 7
		placeholders = append(placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d, $%d, $%d)", base+1, base+2, base+3, base+4, base+5, base+6, base+7),
		)
		coordinates := d.Coordinates()
		args = append(args,
			d.Id(), d.ContractId(), d.Date(), d.Street(), d.Number(), coordinates.Latitude(), coordinates.Longitude(),
		)
	}

	query = fmt.Sprintf(
		"INSERT INTO delivery(id, contract_id, date, street, number, latitude, longitude) VALUES %s RETURNING id, contract_id, date, street, number, latitude, longitude, status",
		strings.Join(placeholders, ","),
	)

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("[repository:contract][Create] delivery bach insert failed: %v", err)
		return nil, fmt.Errorf("delivery batch insert failed: %w", err)
	}
	defer rows.Close()

	var insertedDeliveries []deliveries.Delivery
	for rows.Next() {
		var dId, cId uuid.UUID
		var date time.Time
		var street string
		var number int
		var latitude, longitude float64
		var status deliveries.DeliveryStatus

		if err := rows.Scan(&dId, &cId, &date, &street, &number, &latitude, &longitude, &status); err != nil {
			log.Printf("[repository:contract][Create] scanning rows: %v", err)
			return nil, fmt.Errorf("scanning delivery failed: %w", err)
		}

		coordinates, _ := valueobjects.NewCoordinates(latitude, longitude)
		d := deliveries.NewDeliveryFromDB(dId, cId, date, street, number, *coordinates, status)
		insertedDeliveries = append(insertedDeliveries, *d)
	}

	var contract = contracts.NewContractFromDb(
		id, administratorId, patientId, contractType, contractStatus,
		creation, start, end, cost, insertedDeliveries, createdAt, updatedAt, deletedAt,
	)
	return contract, nil
}

func (r *ContractRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status string) (*contracts.Contract, error) {
	var (
		cId, administratorId, patientId            uuid.UUID
		contractType                               contracts.ContractType
		contractStatus                             contracts.ContractStatus
		creation, start, end, createdAt, updatedAt time.Time
		deletedAt                                  *time.Time
		cost                                       int
	)

	query := `
		UPDATE contract
		SET status = $1
		WHERE id = $2
	`

	err := r.DB.QueryRowContext(
		ctx, query, status, id,
	).Scan(
		&cId, &administratorId, &patientId, &contractType, &contractStatus,
		&creation, &start, &end, &cost, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("[repository:contract][ChangeStatus] error executing SQL query: %v", err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	var contract = contracts.NewContractFromDb(
		id, administratorId, patientId, contractType, contractStatus,
		creation, start, end, cost, nil, createdAt, updatedAt, deletedAt,
	)
	return contract, nil
}

func (r *ContractRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM contract
			WHERE id = $1
		)
	`

	err := r.DB.QueryRowContext(ctx, query, id).Scan(&exist)
	if err != nil {
		log.Printf("[repository:contract][ExistById] error executing SQL query '%s': %v", query, err)
		return false, err
	}

	log.Printf("[repository:contract][ExistsById] id=%s exists=%t", id, exist)
	return exist, nil
}

func (r *ContractRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM contract
	`

	err := r.DB.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Printf("[repository:contract][CountActive] error executing SQL query in Count: %v", err)
		return 0, err
	}
	return count, nil
}

func NewContractRepository(db *sql.DB) contracts.ContractRepository {
	return &ContractRepository{DB: db}
}
