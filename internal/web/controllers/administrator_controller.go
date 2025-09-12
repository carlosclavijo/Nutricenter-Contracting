package controllers

import (
	"database/sql"
	"encoding/json"
	administratorcommands "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command/administrator"
	administratorhandlers "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/handlers/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/service"
	administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type AdministratorHandler struct {
	createHandler *administratorhandlers.CreateAdministratorHandler
	updateHandler *administratorhandlers.UpdateAdministratorHandler
	deleteHandler *administratorhandlers.DeleteAdministratorHandler
}

func NewAdministratorHandler(db *sql.DB) *AdministratorHandler {
	repo := repositories.NewAdministratorRepository(db)
	factory := administrators.NewAdministratorFactory()
	s := service.NewAdministratorService(repo, factory)
	ch := &administratorhandlers.CreateAdministratorHandler{Service: s}
	uh := &administratorhandlers.UpdateAdministratorHandler{Service: s}
	dh := &administratorhandlers.DeleteAdministratorHandler{Service: s}
	return &AdministratorHandler{ch, uh, dh}
}

func (h *AdministratorHandler) CreateAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	cmd := administratorcommands.CreateAdministratorCommand{
		Name:  req.Name,
		Phone: req.Phone,
	}

	admin, err := h.createHandler.Handle(r.Context(), cmd)
	if err != nil {
		log.Printf("failed to create administrator: %v", err)
		http.Error(w, "failed to create an administrator", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdministratorHandler) UpdateAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(req.Id)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
	}

	cmd := administratorcommands.UpdateAdministratorCommand{
		Id:    uid,
		Name:  req.Name,
		Phone: req.Phone,
	}

	admin, err := h.updateHandler.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, "failed to update the administrator", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdministratorHandler) DeleteAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(req.Id)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
	}

	admin, err := h.deleteHandler.Handle(r.Context(), uid)
	if err != nil {
		http.Error(w, "failed to delete the administrator", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdministratorHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.CreateAdministrator)
	r.Put("/", h.UpdateAdministrator)
	r.Delete("/", h.DeleteAdministrator)
}
