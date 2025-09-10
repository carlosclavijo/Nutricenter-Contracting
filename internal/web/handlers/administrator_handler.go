package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/handlers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/service"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrators"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AdministratorHandler struct {
	handler *handlers.CreateAdministratorHandler
}

func NewAdministratorHandler(db *sql.DB) *AdministratorHandler {
	repo := repositories.NewAdministratorRepository(db)
	factory := administrators.NewAdministratorFactory()
	s := service.NewAdministratorService(repo, factory)
	h := &handlers.CreateAdministratorHandler{Service: s}
	return &AdministratorHandler{h}
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

	cmd := command.CreateAdministratorCommand{
		Name:  req.Name,
		Phone: req.Phone,
	}

	id, err := h.handler.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, "failed to crate an administrator", http.StatusInternalServerError)
	}
	resp := map[string]string{"id": id.String()}
	json.NewEncoder(w).Encode(resp)
}

func (h *AdministratorHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.CreateAdministrator)
}
