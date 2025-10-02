package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	command "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/handlers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	query "github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/handlers/contract"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type contractFull struct {
	Id              uuid.UUID
	AdministratorId uuid.UUID
	PatientId       uuid.UUID
	ContractType    contracts.ContractType
	ContractStatus  contracts.ContractStatus
	CreationDate    time.Time
	StartDate       time.Time
	EndDate         time.Time
	CostValue       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	Deliveries      []deliveryFull
}

type deliveryFull struct {
	Id         uuid.UUID
	ContractId uuid.UUID
	Date       time.Time
	Street     string
	Number     int
	Latitude   float64
	Longitude  float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type ContractController struct {
	cmdHandler command.ContractHandler
	qryHandler query.ContractHandler
}

func NewContractHandler(db *sql.DB) *ContractController {
	repo := repositories.NewContractRepository(db)
	rAdm := repositories.NewAdministratorRepository(db)
	rPtn := repositories.NewPatientRepository(db)
	factory := contracts.NewContractFactory()
	cmdHandler := command.NewContractHandler(repo, factory)
	qryHandler := query.NewContractHandler(repo, rAdm, rPtn, factory)
	return &ContractController{*cmdHandler, *qryHandler}
}

func (h *ContractController) GetAllContracts(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetAllContractsQuery{}
	contractlist, err := h.qryHandler.HandleGetAll(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:contract][GetAllContracts] failed to fetch contract: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: "Could not fetch contracts",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[[]*dto.ContractDTO]{
		Success: true,
		Data:    contractlist,
		Length:  len(contractlist),
	})
}

func (h *ContractController) GetContractById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:contract][GetContractById] invalid UUID format '%s': %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_ID_FORMAT",
				Message: "The provided ID is not a valid UUID",
			},
		})
		return
	}

	qry := queries.GetContractByIdQuery{Id: id}
	cntrct, err := h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:contract][GetContractById] failed to fetch contract by id '%s': %v", idStr, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not fetch contract by ID",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[dto.ContractDTO]{
		Success: true,
		Data:    *cntrct,
	})
}

func (h *ContractController) CreateContract(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AdministratorId uuid.UUID `json:"administrator_id"`
		PatientId       uuid.UUID `json:"patient_id"`
		ContractType    string    `json:"contract_type"`
		Start           time.Time `json:"start"`
		Cost            int       `json:"cost"`
		Street          string    `json:"street"`
		Number          int       `json:"number"`
		Latitude        float64   `json:"latitude"`
		Longitude       float64   `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:contract][CreateContract] failed to decode request body '%v': %v", req, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	cmd := commands.CreateContractCommand{
		AdministratorId: req.AdministratorId,
		PatientId:       req.PatientId,
		ContractType:    req.ContractType,
		StartDate:       req.Start,
		Cost:            req.Cost,
		Street:          req.Street,
		Number:          req.Number,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
	}

	cntrct, err := h.cmdHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:contract][CreateContract] failed to create contract with command '%v': %v", cntrct, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create contract",
			},
		})
		return
	}

	cntFull := mapToContractFull(cntrct)
	writeJSON(w, http.StatusCreated, helpers.Response[contractFull]{
		Success: true,
		Data:    cntFull,
	})
}

func (h *ContractController) ChangeStatusContract(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id     uuid.UUID `json:"id"`
		Status string    `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:contract][ChangeStatusContract] failed to decode request body '%v': %v", req, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	cmd := commands.ChangeStatusContractCommand{
		Id:     req.Id,
		Status: req.Status,
	}

	cntrct, err := h.cmdHandler.HandleChangeStatus(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:contract][ChangeStatusContract] failed to change contract status with command '%v': %v", cntrct, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create contract",
			},
		})
		return
	}

	cntFull := mapToContractFull(cntrct)
	writeJSON(w, http.StatusCreated, helpers.Response[contractFull]{
		Success: true,
		Data:    cntFull,
	})
}

func mapToContractFull(c *contracts.Contract) contractFull {
	var deliveries []deliveryFull
	for _, v := range c.Deliveries() {
		c := v.Coordinates()
		d := deliveryFull{
			Id:         v.Id(),
			ContractId: v.ContractId(),
			Date:       v.Date(),
			Street:     v.Street(),
			Number:     v.Number(),
			Latitude:   c.Latitude(),
			Longitude:  c.Longitude(),
			CreatedAt:  v.CreatedAt(),
			UpdatedAt:  v.UpdatedAt(),
			DeletedAt:  v.DeletedAt(),
		}
		deliveries = append(deliveries, d)
	}
	return contractFull{
		Id:              c.Id(),
		AdministratorId: c.AdministratorId(),
		PatientId:       c.PatientId(),
		ContractType:    c.ContractType(),
		ContractStatus:  c.ContractStatus(),
		CreationDate:    c.CreationDate(),
		StartDate:       c.StartDate(),
		EndDate:         c.EndDate(),
		CostValue:       c.CostValue(),
		CreatedAt:       c.CreatedAt(),
		UpdatedAt:       c.UpdatedAt(),
		DeletedAt:       c.DeletedAt(),
		Deliveries:      deliveries,
	}
}

func (h *ContractController) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllContracts)
	r.Get("/{id}", h.GetAllContracts)
	r.Post("/", h.CreateContract)
	r.Post("/status", h.ChangeStatusContract)
}
