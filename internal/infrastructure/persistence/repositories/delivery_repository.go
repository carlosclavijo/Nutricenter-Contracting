package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/google/uuid"
	"log"
)

type DeliveryRepository struct {
	DB *sql.DB
}

func (d DeliveryRepository) GetAll(ctx context.Context) (*[]dto.DeliveryDTO, error) {
	var deliveries []dto.DeliveryDTO

	query := `
		SELECT id, contract_id, date, street, number, latitude, longitude, status
		FROM delivery
	`

	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[repository:delivery][GetAll] Error executing query '%s': %v", query, err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("[repository:delivery][GetAll] failed to close rows: %v", err)
			return
		}
	}(rows)
	for rows.Next() {
		var d dto.DeliveryDTO
		err = rows.Scan(&d.Id, &d.ContractId, &d.Date, &d.Street, &d.Number, &d.Latitude, &d.Longitude, &d.Status)
		if err != nil {
			log.Printf("[repository:delivery][GetAll] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		deliveries = append(deliveries, d)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:delivery][GetAll] error reading deliveries: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	log.Printf("[repository:delivery][GetAll] successfully fetched %d deliveries", len(deliveries))
	return &deliveries, nil
}

func (d DeliveryRepository) GetById(ctx context.Context, id uuid.UUID) (*dto.DeliveryDTO, error) {
	var delivery dto.DeliveryDTO

	query := `
		SELECT id, contract_id, date, street, number, latitude, longitude, status
		FROM delivery
		WHERE id = $1
	`
	err := d.DB.QueryRowContext(ctx, query, id).Scan(&delivery.Id, &delivery.ContractId, &delivery.Date, &delivery.Street, &delivery.Number, &delivery.Latitude, &delivery.Longitude, &delivery.Status)
	if err != nil {
		log.Printf("[repository:delivery][GetById] error executing SQL query '%s': %v", query, err)
		return nil, fmt.Errorf("query failed: %w", err)
	}

	log.Printf("[repository:velivery][GetById] successfully fetched verlivery")
	return &delivery, nil
}

func (d DeliveryRepository) GetListByContractId(ctx context.Context, contractId uuid.UUID) (*[]dto.DeliveryDTO, error) {
	var deliveries []dto.DeliveryDTO

	query := `
		SELECT id, contract_id, date, street, number, latitude, longitude, status
		FROM delivery
		WHERE contract_id = $1
	`

	rows, err := d.DB.QueryContext(ctx, query, contractId)
	if err != nil {
		log.Printf("[repository:delivery][GetListByContractId] Error executing query '%s': %v", query, err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("[repository:delivery][GetListByContractId] failed to close rows: %v", err)
			return
		}
	}(rows)
	for rows.Next() {
		var d dto.DeliveryDTO
		err = rows.Scan(&d.Id, &d.ContractId, &d.Date, &d.Street, &d.Number, &d.Latitude, &d.Longitude, &d.Status)
		if err != nil {
			log.Printf("[repository:delivery][GetListByContractId] error reading adminDTO for a slice of admins: %v", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		deliveries = append(deliveries, d)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[repository:delivery][GetListByContractId] error reading deliveries: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	log.Printf("[repository:delivery][GetListByContractId] successfully fetched %d deliveries", len(deliveries))
	return &deliveries, nil
}

func (d DeliveryRepository) Update(ctx context.Context, contract *deliveries.Delivery) (*deliveries.Delivery, error) {
	//TODO implement me
	panic("implement me")
}

func (d DeliveryRepository) Delete(ctx context.Context, id uuid.UUID) (*deliveries.Delivery, error) {
	//TODO implement me
	panic("implement me")
}

func (d DeliveryRepository) Restore(ctx context.Context, id uuid.UUID) (*deliveries.Delivery, error) {
	//TODO implement me
	panic("implement me")
}

func (d DeliveryRepository) Active(ctx context.Context, id uuid.UUID) (*deliveries.Delivery, error) {
	//TODO implement me
	panic("implement me")
}

func (d DeliveryRepository) Cancel(ctx context.Context, id uuid.UUID) (*deliveries.Delivery, error) {
	//TODO implement me
	panic("implement me")
}
