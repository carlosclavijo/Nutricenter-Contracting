package patients

import (
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	*abstractions.AggregateRoot
	firstName   string
	lastName    string
	email       vo.Email
	password    vo.Password
	gender      vo.Gender
	birth       vo.BirthDate
	phone       *vo.Phone
	lastLoginAt time.Time
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

var (
	ErrEmptyIdPatient            = errors.New("id cannot be nil")
	ErrEmptyFirstNamePatient     = errors.New("first name cannot be empty")
	ErrEmptyLastNamePatient      = errors.New("last name cannot be empty")
	ErrLongFirstNamePatient      = errors.New("first name cannot be longer than 100 characters")
	ErrLongLastNamePatient       = errors.New("last name cannot be longer than 100 characters")
	ErrNonAlphaFirstNamePatient  = errors.New("first name has non alphabetical characters")
	ErrNonAlphaLastNamePatient   = errors.New("last name has non alphabetical characters")
	ErrExistPatient              = errors.New("patient already exist")
	ErrNotFoundPatient           = errors.New("patient not found")
	ErrInvalidCredentialsPatient = errors.New("invalid credentials")
)

func NewPatient(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) *Patient {
	return &Patient{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		password:      password,
		gender:        gender,
		birth:         birth,
		phone:         phone,
	}
}

func (p *Patient) Id() uuid.UUID {
	return p.Entity.Id
}

func (p *Patient) FirstName() string {
	return p.firstName
}

func (p *Patient) LastName() string {
	return p.lastName
}

func (p *Patient) Email() vo.Email {
	return p.email
}

func (p *Patient) Password() vo.Password {
	return p.password
}

func (p *Patient) Gender() vo.Gender {
	return p.gender
}

func (p *Patient) Birth() vo.BirthDate {
	return p.birth
}

func (p *Patient) Phone() *vo.Phone {
	return p.phone
}

func (p *Patient) LastLoginAt() time.Time {
	return p.lastLoginAt
}

func (p *Patient) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Patient) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Patient) DeletedAt() *time.Time {
	return p.deletedAt
}

func (p *Patient) Logged() {
	p.lastLoginAt = time.Now()
}

func NewPatientFromDB(id uuid.UUID, firstName, lastName, email, password, gender string, birth time.Time, phone *string, lastLoginAt, createdAt, updatedAt time.Time, deletedAt *time.Time) (*Patient, error) {
	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	passwordVo, err := vo.NewHashedPassword(password)
	if err != nil {
		return nil, err
	}

	genderVo, err := vo.ParseGender(gender)
	if err != nil {
		return nil, err
	}

	birthVo, err := vo.NewBirthDate(birth)
	if err != nil {
		return nil, err
	}

	phoneVo, err := vo.NewPhone(phone)
	if err != nil {
		return nil, err
	}

	return &Patient{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		firstName:     firstName,
		lastName:      lastName,
		email:         emailVo,
		password:      passwordVo,
		gender:        genderVo,
		birth:         birthVo,
		phone:         phoneVo,
		lastLoginAt:   lastLoginAt,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}, nil
}
