package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/command/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	command "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/handlers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	query "github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/handlers/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	Password  string     `json:"password,omitempty"`
	Gender    string     `json:"gender"`
	Birth     *time.Time `json:"birth,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	LastLogin time.Time  `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (h *AdministratorController) GetAllAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := administratorqueries.GetAllAdministratorsQuery{}
	admins, err := h.qryHandler.HandleGetAll(r.Context(), qry)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("[controller:administrator][GetAllAdministrators] failed to fetch administrators: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: "Could not fetch administrators",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][GetAllAdministrators] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[[]dto.AdministratorDTO]{
		Success: true,
		Data:    *admins,
		Length:  len(*admins),
	}); err != nil {
		log.Printf("[controller:administrator][GetAllAdministrators] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) GetListAdministrators(w http.ResponseWriter, r *http.Request) {
	qry := administratorqueries.GetListAdministratorsQuery{}
	admins, err := h.qryHandler.HandleGetList(r.Context(), qry)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("[controller:administrator][GetListAdministrators] failed to fetch administrators: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIST_FAILED",
				Message: "Could not fetch administrators",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][GetListAdministrators] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[[]dto.AdministratorDTO]{
		Success: true,
		Data:    *admins,
		Length:  len(*admins),
	}); err != nil {
		log.Printf("[controller:administrator][GetListAdministrators] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) GetAdministratorById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorById] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][GetAdministratorById] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to parse UUID error response", http.StatusInternalServerError)
		}
		return
	}

	qry := administratorqueries.GetAdministratorByIdQuery{Id: id}
	admin, err := h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorByIdQuery] failed to retrieve administrator with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve administrator",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][GetAdministratorById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[dto.AdministratorDTO]{
		Success: true,
		Data:    *admin,
	}); err != nil {
		log.Printf("[controller:administrator][GetAdministratorById]  failed to encode administrator with ID %s to JSON: %v", id, err)
		http.Error(w, "Failed to encode administrator data", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) GetAdministratorByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := administratorqueries.GetAdministratorByEmailQuery{Email: email}
	w.Header().Set("Content-Type", "application/json")
	admin, err := h.qryHandler.HandleGetByEmail(r.Context(), qry)

	var birthTm *time.Time
	var phoneStr *string
	if admin.Birth() != nil {
		b := admin.Birth().Value()
		birthTm = b
	}
	if admin.Phone() != nil {
		phoneStr = admin.Phone().String()
	}

	admFull := adminFull{
		Id:        admin.Id(),
		FirstName: admin.FirstName(),
		LastName:  admin.LastName(),
		Email:     admin.Email().Value(),
		Gender:    admin.Gender(),
		Birth:     birthTm,
		Phone:     phoneStr,
		LastLogin: *admin.LastLoginAt(),
		CreatedAt: *admin.CreatedAt(),
		UpdatedAt: *admin.UpdatedAt(),
		DeletedAt: admin.DeletedAt(),
	}

	if err != nil {
		log.Printf("[controller:administrator][GetAdministratorByEmail]  failed to retrieve administrator with Email '%s': %v", email, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_EMAIL_FAILED",
				Message: "Could not retrieve administrator",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][GetAdministratorByEmail]failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:administrator][GetAdministratorById] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) ExistAdministratorById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorById] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][ExistAdministratorById] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	qry := administratorqueries.ExistAdministratorByIdQuery{Id: id}
	exist, err := h.qryHandler.HandleExistById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorById] failed to retrieve if the administrator exist or not with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXISTS_BY_ID_FAILED",
				Message: "Could not retrieve if administrator exist or not",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][ExistAdministratorById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	if err = json.NewEncoder(w).Encode(helpers.Response[bool]{
		Success: true,
		Data:    exist,
	}); err != nil {
		log.Printf("[controller:administrator][ExistAdministratorById] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) ExistAdministratorByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := administratorqueries.ExistAdministratorByEmailQuery{Email: email}
	w.Header().Set("Content-Type", "application/json")

	exist, err := h.qryHandler.HandleExistByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][ExistAdministratorByEmail] failed to retrieve if the administrator exist or not with email %s: %v", email, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXIST_BY_EMAIL_FAILED",
				Message: "Could not retrieve if administrator exist or not",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][ExistAdministratorById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[bool]{
		Success: true,
		Data:    exist,
	}); err != nil {
		log.Printf("[controller:administrator][ExistAdministratorByEmail] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:administrator][Login] failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][Login] failed to encode request body response: %v", encErr)
			http.Error(w, "Failed to encode request body response", http.StatusInternalServerError)
		}
		return
	}

	qry := administratorqueries.GetAdministratorByEmailQuery{
		Email: req.Email,
	}

	admin, err := h.qryHandler.HandleGetByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:administrator][Login] failed to retrieve administrator with Email '%s': %v", req.Email, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOGIN_FAILED",
				Message: "Could not retrieve administrator",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][Login] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("[controller:administrator][Login] failed to hash password %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "HASHING_PASSWORD_FAILED",
				Message: "Could not hash password",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][Login] failed to encode hash password response: %v", encErr)
			http.Error(w, "failed to encode hash password response", http.StatusInternalServerError)
		}
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil || req.Email != admin.Email().Value() {
		log.Printf("[controller:administrator][Login] login failed, invalid credentials for email=%s", req.Email)
		w.WriteHeader(http.StatusUnauthorized)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COMPARING_HASHED_PASSWORD_FAILED",
				Message: "Invalid credentials for email",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][Login] failed to encode hash comparison response: %v", encErr)
			http.Error(w, "failed to encode hash comparison response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[administrators.Administrator]{
		Success: true,
		Data:    *admin,
	}); err != nil {
		log.Printf("[controller:administrator][Login] failed to encode administrator with email '%s' to JSON: %v", req.Email, err)
		http.Error(w, "Failed to encode administrator data", http.StatusInternalServerError)
	}
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
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to decode request body '%v': %v", req, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][CreateAdministrator] failed to encode request body response: %v", encErr)
			http.Error(w, "Failed to encode request body response", http.StatusInternalServerError)
		}
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to hash password %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "HASHING_PASSWORD_FAILED",
				Message: "Could not hash password",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][CreateAdministrator] failed to hash password response: %v", encErr)
			http.Error(w, "Failed to encode hash password response", http.StatusInternalServerError)
		}
		return
	}

	cmd := administratorcommands.CreateAdministratorCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Gender:    req.Gender,
		Birth:     req.Birth,
		Phone:     req.Phone,
	}

	admin, err := h.cmdHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to create administrator with command '%v': %v", admin, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create administrator",
			},
		}); encErr != nil {
			log.Printf("[controller:administrator][CreateAdministrator] failed to encode error response: %v", encErr)
			http.Error(w, "failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	log.Print("[controller:administrator][CreateAdministrator] successfully created administrator")

	var birthTm *time.Time
	var phoneStr *string
	if admin.Birth() != nil {
		b := admin.Birth().Value()
		birthTm = b
	}
	if admin.Phone() != nil {
		phoneStr = admin.Phone().String()
	}

	admFull := adminFull{
		Id:        admin.Id(),
		FirstName: admin.FirstName(),
		LastName:  admin.LastName(),
		Email:     admin.Email().Value(),
		Gender:    admin.Gender(),
		Birth:     birthTm,
		Phone:     phoneStr,
		LastLogin: *admin.LastLoginAt(),
		CreatedAt: *admin.CreatedAt(),
		UpdatedAt: *admin.UpdatedAt(),
		DeletedAt: admin.DeletedAt(),
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:administrator][CreateAdministrator] failed to encode success response: %v", err)
		http.Error(w, "failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) UpdateAdministrator(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id        string     `json:"id"`
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		Birth     *time.Time `json:"birth"`
		Phone     *string    `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("web/controllers: failed to decode request body in UpdateAdministrator: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("web/controllers: invalid UUID in UpdateAdministrator: %q, error: %v", req.Id, err)
		http.Error(w, "Invalid administrator ID format", http.StatusBadRequest)
		return
	}

	cmd := administratorcommands.UpdateAdministratorCommand{
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
		log.Printf("web/controllers: failed to update administrator with ID %s: %v", uid, err)
		http.Error(w, "Failed to update administrator", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(admin); err != nil {
		log.Printf("web/controllers: failed to encode updated administrator with ID %s to JSON: %v", uid, err)
		http.Error(w, "Failed to encode administrator data", http.StatusInternalServerError)
	}
}

func (h *AdministratorController) DeleteAdministrator(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("web/controllers: invalid UUID in DeleteAdministrator: %q, error: %v", idStr, err)
		http.Error(w, "Invalid administrator ID format", http.StatusBadRequest)
		return
	}

	if _, err = h.qryHandler.HandleGetById(r.Context(), administratorqueries.GetAdministratorByIdQuery{Id: id}); err != nil {
		log.Printf("web/controllers: administrator with ID %s not found in DeleteAdministrator: %v", id, err)
		http.Error(w, "Administrator not found", http.StatusNotFound)
		return
	}

	admin, err := h.cmdHandler.HandleDelete(r.Context(), id)
	if err != nil {
		log.Printf("web/controllers: failed to delete administrator with ID %s: %v", id, err)
		http.Error(w, "Failed to delete administrator", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(admin); err != nil {
		log.Printf("web/controllers: failed to encode deleted administrator with ID %s to JSON: %v", id, err)
		http.Error(w, "Failed to encode administrator data", http.StatusInternalServerError)
	}
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (h *AdministratorController) RegisterRoutes(r chi.Router) {
	r.Get("/all", h.GetAllAdministrators)
	r.Get("/list", h.GetListAdministrators)
	r.Get("/email/{email}", h.GetAdministratorByEmail)
	r.Get("/exist/id/{id}", h.ExistAdministratorById)
	r.Get("/exist/email/{email}", h.ExistAdministratorByEmail)
	r.Get("/{id}", h.GetAdministratorById)
	r.Post("/", h.CreateAdministrator)
	r.Post("/login", h.Login)
	r.Put("/", h.UpdateAdministrator)
	r.Delete("/{id}", h.DeleteAdministrator)
}
