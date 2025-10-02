package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	command "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/handlers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	query "github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/handlers/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type AdministratorController struct {
	cmdHandler command.AdministratorHandler
	qryHandler query.AdministratorHandler
}

func NewAdministratorHandler(db *sql.DB) *AdministratorController {
	repo := repositories.NewAdministratorRepository(db)
	factory := administrators.NewAdministratorFactory()
	cmdHandler := command.NewAdministratorHandler(repo, factory)
	qryHandler := query.NewAdministratorHandler(repo, factory)
	return &AdministratorController{*cmdHandler, *qryHandler}
}

type adminFull struct {
	Id        uuid.UUID  `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Gender    string     `json:"gender"`
	Birth     *time.Time `json:"birth,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	LastLogin time.Time  `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (h *AdministratorController) GetAllAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetAllAdministratorsQuery{}
	admins, err := h.qryHandler.HandleGetAll(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:administrator][GetAllAdministrators] failed to fetch administrators: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: "Could not fetch administrators",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[[]*dto.AdministratorDTO]{
		Success: true,
		Data:    admins,
		Length:  len(admins),
	})
}

func (h *AdministratorController) GetListAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetListAdministratorsQuery{}
	admins, err := h.qryHandler.HandleGetList(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:administrator][GetListAdministrators] failed to fetch administrators: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIST_FAILED",
				Message: "Could not fetch administrators",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[[]*dto.AdministratorDTO]{
		Success: true,
		Data:    admins,
		Length:  len(admins),
	})
}

func (h *AdministratorController) GetAdministratorById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorById] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	qry := queries.GetAdministratorByIdQuery{Id: id}
	admin, err := h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorById] failed to retrieve administrator with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve administrator",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[dto.AdministratorDTO]{
		Success: true,
		Data:    *admin,
	})
}

func (h *AdministratorController) GetAdministratorByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := queries.GetAdministratorByEmailQuery{Email: email}
	admin, err := h.qryHandler.HandleGetByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorByEmail] failed to retrieve administrator with Email '%s': %v", email, err)
		_ = json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_EMAIL_FAILED",
				Message: "Could not retrieve administrator",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[dto.AdministratorDTO]{
		Success: true,
		Data:    *admin,
	})
}

func (h *AdministratorController) ExistAdministratorById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorById] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	qry := queries.ExistAdministratorByIdQuery{Id: id}
	exist, err := h.qryHandler.HandleExistById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorById] failed to retrieve if the administrator exists with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXISTS_BY_ID_FAILED",
				Message: "Could not retrieve if administrator exists or not",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[bool]{
		Success: true,
		Data:    exist,
	})
}

func (h *AdministratorController) ExistAdministratorByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := queries.ExistAdministratorByEmailQuery{Email: email}

	exist, err := h.qryHandler.HandleExistByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorByEmail] failed to retrieve if the administrator exists with email %s: %v", email, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXIST_BY_EMAIL_FAILED",
				Message: "Could not retrieve if administrator exists or not",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[bool]{
		Success: true,
		Data:    exist,
	})
}

func (h *AdministratorController) LoginAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:administrator][Login] failed to decode request body: %v", err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	qry := commands.LoginAdministratorCommand{
		Email:    req.Email,
		Password: req.Password,
	}

	admin, err := h.cmdHandler.HandleLogin(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][Login] failed to retrieve administrator with Email '%s': %v", req.Email, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOGIN_FAILED",
				Message: "Could not retrieve administrator",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[administrators.Administrator]{
		Success: true,
		Data:    *admin,
	})
}

func (h *AdministratorController) CreateAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		Gender    string     `json:"gender"`
		Birth     *time.Time `json:"birth"`
		Phone     *string    `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to decode request body '%v': %v", req, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	cmd := commands.CreateAdministratorCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Gender:    req.Gender,
		Birth:     req.Birth,
		Phone:     req.Phone,
	}

	admin, err := h.cmdHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to create administrator with command '%v': %v", admin, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create administrator",
			},
		})
		return
	}

	admFull := mapToAdminFull(admin)
	writeJSON(w, http.StatusCreated, helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	})
}

func (h *AdministratorController) UpdateAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id        string     `json:"id"`
		FirstName string     `json:"first_name,omitempty"`
		LastName  string     `json:"last_name,omitempty"`
		Email     string     `json:"email,omitempty"`
		Password  string     `json:"password,omitempty"`
		Gender    string     `json:"gender,omitempty"`
		Birth     *time.Time `json:"birth,omitempty"`
		Phone     *string    `json:"phone,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:administrator][UpdateAdministrator] failed to decode request body: %v", err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	uid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("[controller:administrator][UpdateAdministrator] invalid UUID '%s', error: %v", req.Id, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_UUID_PARSING",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	cmd := commands.UpdateAdministratorCommand{
		Id:        uid,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Birth:     req.Birth,
		Phone:     req.Phone,
	}

	admin, err := h.cmdHandler.HandleUpdate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:administrator][UpdateAdministrator] failed to update administrator with ID %s: %v", uid, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPDATE_FAILED",
				Message: "Could not update administrator",
			},
		})
		return
	}

	admFull := mapToAdminFull(admin)
	writeJSON(w, http.StatusOK, helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	})
}

func (h *AdministratorController) DeleteAdministrator(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:administrator][DeleteAdministrator] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	admin, err := h.cmdHandler.HandleDelete(r.Context(), id)
	if err != nil {
		log.Printf("[controller:administrator][DeleteAdministrator] failed to delete administrator with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "DELETE_FAILED",
				Message: "Could not delete administrator",
			},
		})
		return
	}

	admFull := mapToAdminFull(admin)
	writeJSON(w, http.StatusOK, helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	})
}

func (h *AdministratorController) RestoreAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:administrator][RestoreAdministrator] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	admin, err := h.cmdHandler.HandleRestore(r.Context(), id)
	if err != nil {
		log.Printf("[controller:administrator][RestoreAdministrator] failed to restore administrator with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "RESTORE_FAILED",
				Message: "Could not restore administrator",
			},
		})
		return
	}

	admFull := mapToAdminFull(admin)
	writeJSON(w, http.StatusOK, helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	})
}

func (h *AdministratorController) CountAllAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountAllAdministratorsQuery{}
	count, err := h.qryHandler.HandleCountAll(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][CountAllAdministrators] failed to get quantity: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_ALL_FAILED",
				Message: "Could not get quantity",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[int]{
		Success: true,
		Length:  count,
	})
}

func (h *AdministratorController) CountActiveAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountActiveAdministratorsQuery{}
	count, err := h.qryHandler.HandleCountActive(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][CountActiveAdministrators] failed to get quantity: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_ACTIVE_FAILED",
				Message: "Could not get quantity",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[int]{
		Success: true,
		Length:  count,
	})
}

func (h *AdministratorController) CountDeletedAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountDeletedAdministratorsQuery{}
	count, err := h.qryHandler.HandleCountDeleted(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][CountDeletedAdministrators] failed to get quantity: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_DELETED_FAILED",
				Message: "Could not get quantity",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[int]{
		Success: true,
		Length:  count,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func mapToAdminFull(admin *administrators.Administrator) adminFull {
	var birthTm *time.Time
	if admin.Birth() != nil {
		b := admin.Birth().Value()
		birthTm = b
	}

	var phoneStr *string
	if admin.Phone() != nil {
		p := admin.Phone().String()
		phoneStr = p
	}

	return adminFull{
		Id:        admin.Id(),
		FirstName: admin.FirstName(),
		LastName:  admin.LastName(),
		Email:     admin.Email().Value(),
		Password:  "",
		Gender:    admin.Gender().String(),
		Birth:     birthTm,
		Phone:     phoneStr,
		LastLogin: admin.LastLoginAt,
		CreatedAt: admin.CreatedAt(),
		UpdatedAt: admin.UpdatedAt,
		DeletedAt: admin.DeletedAt,
	}
}

func (h *AdministratorController) RegisterRoutes(r chi.Router) {
	r.Get("/all", h.GetAllAdministrators)
	r.Get("/list", h.GetListAdministrators)
	r.Get("/email/{email}", h.GetAdministratorByEmail)
	r.Get("/exist/id/{id}", h.ExistAdministratorById)
	r.Get("/exist/email/{email}", h.ExistAdministratorByEmail)
	r.Get("/count/all", h.CountAllAdministrators)
	r.Get("/count/active", h.CountActiveAdministrators)
	r.Get("/count/deleted", h.CountDeletedAdministrators)
	r.Get("/{id}", h.GetAdministratorById)
	r.Post("/", h.CreateAdministrator)
	r.Post("/login", h.LoginAdministrator)
	r.Put("/", h.UpdateAdministrator)
	r.Patch("/", h.RestoreAdministrator)
	r.Delete("/{id}", h.DeleteAdministrator)
}
