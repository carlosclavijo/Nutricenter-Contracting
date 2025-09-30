package controllers

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

type patientFull struct {
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
	qry := queries.GetAllPatientsQuery{}
	ptnts, err := h.qryHandler.HandleGetAll(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][GetAllPatients] failed to fetch patients: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: "Could not fetch patients",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[[]dto.PatientDTO]{
		Success: true,
		Data:    *ptnts,
		Length:  len(*ptnts),
	})
}

func (h *PatientController) GetListPatients(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetListPatientsQuery{}
	ptnts, err := h.qryHandler.HandleGetList(r.Context(), qry)

	if err != nil {
		log.Printf("[controller:patient][GetListPatients] failed to fetch patients: %v", err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIST_FAILED",
				Message: "Could not fetch patients",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[[]dto.PatientDTO]{
		Success: true,
		Data:    *ptnts,
		Length:  len(*ptnts),
	})
}

func (h *PatientController) GetPatientById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Printf("[controller:patient][GetPatientById] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	qry := queries.GetPatientByIdQuery{Id: id}
	ptnt, err := h.qryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][GetPatientById] failed to retrieve patient with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: "Could not retrieve patient",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[dto.PatientDTO]{
		Success: true,
		Data:    *ptnt,
	})
}

func (h *PatientController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := queries.GetPatientByEmailQuery{Email: email}
	ptnt, err := h.qryHandler.HandleGetByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][GetPatientByEmail] failed to retrieve patient with Email '%s': %v", email, err)
		_ = json.NewEncoder(w).Encode(helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_EMAIL_FAILED",
				Message: "Could not retrieve patient",
			},
		})
		return
	}

	ptnFull := mapToPatientFull(ptnt)
	writeJSON(w, http.StatusOK, helpers.Response[patientFull]{
		Success: true,
		Data:    ptnFull,
	})
}

func (h *PatientController) ExistPatientById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientById] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	qry := queries.ExistPatientByIdQuery{Id: id}
	exist, err := h.qryHandler.HandleExistById(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientById] failed to retrieve if the patient exists with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXISTS_BY_ID_FAILED",
				Message: "Could not retrieve if patient exists or not",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[bool]{
		Success: true,
		Data:    exist,
	})
}

func (h *PatientController) ExistPatientByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	qry := queries.ExistPatientByEmailQuery{Email: email}

	exist, err := h.qryHandler.HandleExistByEmail(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][ExistPatientByEmail] failed to retrieve if the patient exists with email %s: %v", email, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "EXIST_BY_EMAIL_FAILED",
				Message: "Could not retrieve if patient exists or not",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[bool]{
		Success: true,
		Data:    exist,
	})
}

func (h *PatientController) LoginPatient(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[controller:patient][Login] failed to decode request body: %v", err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	qry := commands.LoginPatientCommand{
		Email:    req.Email,
		Password: req.Password,
	}

	ptnt, err := h.cmdHandler.HandleLogin(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][Login] failed to retrieve patient with Email '%s': %v", req.Email, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOGIN_FAILED",
				Message: "Could not retrieve patient",
			},
		})
		return
	}

	writeJSON(w, http.StatusOK, helpers.Response[patients.Patient]{
		Success: true,
		Data:    *ptnt,
	})
}

func (h *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
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
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON format or fields",
			},
		})
		return
	}

	cmd := commands.CreatePatientCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Gender:    req.Gender,
		Birth:     req.Birth,
		Phone:     req.Phone,
	}

	ptnt, err := h.cmdHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:patient][CreatePatient] failed to create patient with command '%v': %v", ptnt, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create patient",
			},
		})
		return
	}

	ptntFull := mapToPatientFull(ptnt)
	writeJSON(w, http.StatusCreated, helpers.Response[patientFull]{
		Success: true,
		Data:    ptntFull,
	})
}

func (h *PatientController) UpdatePatient(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("[controller:patient][UpdatePatient] failed to decode request body: %v", err)
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
		log.Printf("[controller:patient][UpdatePatient] invalid UUID '%s', error: %v", req.Id, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_UUID_PARSING",
				Message: "Could not parse UUID",
			},
		})
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

	ptnt, err := h.cmdHandler.HandleUpdate(r.Context(), cmd)
	if err != nil {
		log.Printf("[controller:patient][UpdatePatient] failed to update patient with ID %s: %v", uid, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPDATE_FAILED",
				Message: "Could not update patient",
			},
		})
		return
	}

	ptntFull := mapToPatientFull(ptnt)
	writeJSON(w, http.StatusOK, helpers.Response[patientFull]{
		Success: true,
		Data:    ptntFull,
	})
}

func (h *PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:patient][DeletePatient] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	ptnt, err := h.cmdHandler.HandleDelete(r.Context(), id)
	if err != nil {
		log.Printf("[controller:patient][DeletePatient] failed to delete patient with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "DELETE_FAILED",
				Message: "Could not delete patient",
			},
		})
		return
	}

	pntnFull := mapToPatientFull(ptnt)
	writeJSON(w, http.StatusOK, helpers.Response[patientFull]{
		Success: true,
		Data:    pntnFull,
	})
}

func (h *PatientController) RestorePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("[controller:patient][RestorePatient] invalid UUID: %q, error: %v", idStr, err)
		writeJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: "Could not parse UUID",
			},
		})
		return
	}

	ptnt, err := h.cmdHandler.HandleRestore(r.Context(), id)
	if err != nil {
		log.Printf("[controller:patient][RestorePatient] failed to restore patient with ID %s: %v", id, err)
		writeJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "RESTORE_FAILED",
				Message: "Could not restore patient",
			},
		})
		return
	}

	pntnFull := mapToPatientFull(ptnt)
	writeJSON(w, http.StatusOK, helpers.Response[patientFull]{
		Success: true,
		Data:    pntnFull,
	})
}

func (h *PatientController) CountAllPatients(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountAllPatientsQuery{}
	count, err := h.qryHandler.HandleCountAll(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][CountAllPatients] failed to get quantity: %v", err)
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

func (h *PatientController) CountActivePatients(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountActivePatientsQuery{}
	count, err := h.qryHandler.HandleCountActive(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][CountActivePatients] failed to get quantity: %v", err)
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

func (h *PatientController) CountDeletedPatients(w http.ResponseWriter, r *http.Request) {
	qry := queries.CountDeletedPatientsQuery{}
	count, err := h.qryHandler.HandleCountDeleted(r.Context(), qry)
	if err != nil {
		log.Printf("[controller:patient][CountDeletedPatients] failed to get quantity: %v", err)
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

func mapToPatientFull(admin *patients.Patient) patientFull {
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

	return patientFull{
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
