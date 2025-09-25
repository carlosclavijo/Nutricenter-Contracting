package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	command "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/handlers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	query "github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/handlers/patient"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type PatientController struct {
	cmdHandler command.PatientHandler
	qryHandler query.PatientHandler
}

func NewPatientHandler(db *sql.DB) *PatientController {
	repo := repositories.NewPatientRepository(db)
	factory := patients.NewPatientFactory()
	cmdHandler := command.NewPatientHandler(repo, factory)
	qryHandler := query.NewPatientHandler(repo, factory)
	return &PatientController{*cmdHandler, *qryHandler}
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

func (h *PatientController) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qry := queries.GetAllPatientsQuery{}
	admins, err := h.qryHandler.HandleGetAll(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][GetAllPatients] failed to fetch patients: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: "Could not fetch patients",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][GetAllPatients] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[[]dto.PatientDTO]{
		Success: true,
		Data:    *admins,
		Length:  len(*admins),
	}); err != nil {
		log.Printf("[controller:patient][GetAllPatients] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) GetListPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qry := queries.GetListPatientsQuery{}
	admins, err := h.qryHandler.HandleGetList(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][GetListPatients] failed to fetch patients: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIST_FAILED",
				Message: "Could not fetch patients",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][GetListPatients] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[[]dto.PatientDTO]{
		Success: true,
		Data:    *admins,
		Length:  len(*admins),
	}); err != nil {
		log.Printf("[controller:patient][GetListPatients] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) GetPatientById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Printf("[controller:patient][GetPatientById] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][GetPatientById] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to parse UUID error response", http.StatusInternalServerError)
		}
		return
	}

	qry := queries.GetPatientByIdQuery{Id: id}
	admin, err := h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][GetPatientById] failed to retrieve patient with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][GetPatientById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[dto.PatientDTO]{
		Success: true,
		Data:    *admin,
	}); err != nil {
		log.Printf("[controller:patient][GetPatientById]  failed to encode patient with ID %s to JSON: %v", id, err)
		http.Error(w, "Failed to encode patient data", http.StatusInternalServerError)
	}
}

func (h *PatientController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := chi.URLParam(r, "email")
	qry := queries.GetPatientByEmailQuery{Email: email}
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
		log.Printf("[controller:patient][GetPatientByEmail]  failed to retrieve patient with Email '%s': %v", email, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_EMAIL_FAILED",
				Message: "Could not retrieve patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][GetPatientByEmail]failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:patient][GetPatientById] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) ExistPatientById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientById] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][ExistPatientById] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	qry := queries.ExistPatientByIdQuery{Id: id}
	exist, err := h.qryHandler.HandleExistById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientById] failed to retrieve if the patient exist or not with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXISTS_BY_ID_FAILED",
				Message: "Could not retrieve if patient exist or not",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][ExistPatientById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	if err = json.NewEncoder(w).Encode(helpers.Response[bool]{
		Success: true,
		Data:    exist,
	}); err != nil {
		log.Printf("[controller:patient][ExistPatientById] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) ExistPatientByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := chi.URLParam(r, "email")
	qry := queries.ExistPatientByEmailQuery{Email: email}

	exist, err := h.qryHandler.HandleExistByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientByEmail] failed to retrieve if the patient exist or not with email %s: %v", email, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXIST_BY_EMAIL_FAILED",
				Message: "Could not retrieve if patient exist or not",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][ExistPatientById] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[bool]{
		Success: true,
		Data:    exist,
	}); err != nil {
		log.Printf("[controller:patient][ExistPatientByEmail] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) LoginPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:patient][Login] failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][Login] failed to encode request body response: %v", encErr)
			http.Error(w, "Failed to encode request body response", http.StatusInternalServerError)
		}
		return
	}

	qry := queries.GetPatientByEmailQuery{
		Email: req.Email,
	}

	admin, err := h.qryHandler.HandleGetByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][Login] failed to retrieve patient with Email '%s': %v", req.Email, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOGIN_FAILED",
				Message: "Could not retrieve patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][Login] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("[controller:patient][Login] failed to hash password %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "HASHING_PASSWORD_FAILED",
				Message: "Could not hash password",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][Login] failed to encode hash password response: %v", encErr)
			http.Error(w, "failed to encode hash password response", http.StatusInternalServerError)
		}
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil || req.Email != admin.Email().Value() {
		log.Printf("[controller:patient][Login] login failed, invalid credentials for email=%s", req.Email)
		w.WriteHeader(http.StatusUnauthorized)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COMPARING_HASHED_PASSWORD_FAILED",
				Message: "Invalid credentials for email",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][Login] failed to encode hash comparison response: %v", encErr)
			http.Error(w, "failed to encode hash comparison response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[patients.Patient]{
		Success: true,
		Data:    *admin,
	}); err != nil {
		log.Printf("[controller:patient][Login] failed to encode patient with email '%s' to JSON: %v", req.Email, err)
		http.Error(w, "Failed to encode patient data", http.StatusInternalServerError)
	}
}

func (h *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
		log.Printf("[controller:patient][CreatePatient] failed to decode request body '%v': %v", req, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CreatePatient] failed to encode request body response: %v", encErr)
			http.Error(w, "Failed to encode request body response", http.StatusInternalServerError)
		}
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("[controller:patient][CreatePatient] failed to hash password %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "HASHING_PASSWORD_FAILED",
				Message: "Could not hash password",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CreatePatient] failed to hash password response: %v", encErr)
			http.Error(w, "Failed to encode hash password response", http.StatusInternalServerError)
		}
		return
	}

	cmd := commands.CreatePatientCommand{
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
		log.Printf("[controller:patient][CreatePatient] failed to create patient with command '%v': %v", admin, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CreatePatient] failed to encode error response: %v", encErr)
			http.Error(w, "failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	log.Print("[controller:patient][CreatePatient] successfully created patient")

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
		log.Printf("[controller:patient][CreatePatient] failed to encode success response: %v", err)
		http.Error(w, "failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		Id        string     `json:"id"`
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		Gender    string     `json:"gender"`
		Birth     *time.Time `json:"birth"`
		Phone     *string    `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:patient][UpdatePatient] failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][UpdatePatient] failed to encode request body response: %v", encErr)
			http.Error(w, "Failed to encode request body response", http.StatusInternalServerError)
		}
		return
	}

	uid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("[controller:patient][UpdatePatient] invalid UUID '%s', error: %v", req.Id, err)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_UUID_PARSING",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][UpdatePatient] failed to parse UUID '%s', error: %v", req.Id, encErr)
			http.Error(w, "Failed to encode hash password response", http.StatusInternalServerError)
		}
		http.Error(w, "Invalid patient ID format", http.StatusBadRequest)
		return
	}

	cmd := commands.UpdatePatientCommand{
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
		log.Printf("[controller:patient][UpdatePatient] failed to update patient with ID %s: %v", uid, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPDATE_FAILED",
				Message: "Could not update patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][UpdatePatient] failed to encode error response: %v", encErr)
			http.Error(w, "failed to encode error response", http.StatusInternalServerError)
		}
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:patient][UpdatePatient] failed to encode success response: %v", err)
		http.Error(w, "failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Printf("[controller:patient][DeletePatient] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][DeletePatient] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to parse UUID error response", http.StatusInternalServerError)
		}
		return
	}

	qry := queries.GetPatientByIdQuery{Id: id}
	_, err = h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][DeletePatient] failed to retrieve patient with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][DeletePatient] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	admin, err := h.cmdHandler.HandleDelete(r.Context(), id)
	if err != nil {
		log.Printf("[controller:patient][DeletePatient] failed to delete patient with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "DELETE_FAILED",
				Message: "Could not delete patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][DeletePatient] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

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

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:patient][DeletePatient] failed to encode success response: %v", err)
		http.Error(w, "failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) RestorePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Printf("[controller:patient][RestorePatient] invalid UUID: %q, error: %v", idStr, err)
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][RestorePatient] failed to parse UUID error response: %v", encErr)
			http.Error(w, "Failed to parse UUID error response", http.StatusInternalServerError)
		}
		return
	}

	qry := queries.GetPatientByIdQuery{Id: id}
	_, err = h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][RestorePatient] failed to retrieve patient with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][RestorePatient] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	admin, err := h.cmdHandler.HandleRestore(r.Context(), id)
	if err != nil {
		log.Printf("[controller:patient][RestorePatient] failed to restore patient with ID %s: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "RESTORE_FAILED",
				Message: "Could not restore patient",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][RestorePatient] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

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

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(helpers.Response[adminFull]{
		Success: true,
		Data:    admFull,
	}); err != nil {
		log.Printf("[controller:patient][RestorePatient] failed to encode success response: %v", err)
		http.Error(w, "failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) CountAllPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qry := queries.CountAllPatientsQuery{}
	count, err := h.qryHandler.HandleCountAll(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][CountAllPatients] failed to get quantity: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_ALL_FAILED",
				Message: "Could not get quantity",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CountAllPatients] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[int]{
		Success: true,
		Length:  count,
	}); err != nil {
		log.Printf("[controller:patient][CountAllPatients] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) CountActivePatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qry := queries.CountActivePatientsQuery{}
	count, err := h.qryHandler.HandleCountActive(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][CountActivePatients] failed to get quantity: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_ACTIVE_FAILED",
				Message: "Could not get quantity",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CountActivePatients] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[int]{
		Success: true,
		Length:  count,
	}); err != nil {
		log.Printf("[controller:patient][CountActivePatients] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func (h *PatientController) CountDeletedPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qry := queries.CountDeletedPatientsQuery{}
	count, err := h.qryHandler.HandleCountDeleted(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][CountDeletedPatients] failed to get quantity: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if encErr := json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "COUNT_DELETED_FAILED",
				Message: "Could not get quantity",
			},
		}); encErr != nil {
			log.Printf("[controller:patient][CountAllPatients] failed to encode error response: %v", encErr)
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Response[int]{
		Success: true,
		Length:  count,
	}); err != nil {
		log.Printf("[controller:patient][CountDeletedPatients] failed to encode success response: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (h *PatientController) RegisterRoutes(r chi.Router) {
	r.Get("/all", h.GetAllPatients)
	r.Get("/list", h.GetListPatients)
	r.Get("/email/{email}", h.GetPatientByEmail)
	r.Get("/exist/id/{id}", h.ExistPatientById)
	r.Get("/exist/email/{email}", h.ExistPatientByEmail)
	r.Get("/count/all", h.CountAllPatients)
	r.Get("/count/active", h.CountActivePatients)
	r.Get("/count/deleted", h.CountDeletedPatients)
	r.Get("/{id}", h.GetPatientById)
	r.Post("/", h.CreatePatient)
	r.Post("/login", h.LoginPatient)
	r.Put("/", h.UpdatePatient)
	r.Patch("/", h.RestorePatient)
	r.Delete("/{id}", h.DeletePatient)
}
