package administrators

import (
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"time"
)

type Administrator struct {
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
	ErrEmptyIdAdministrator            = errors.New("id cannot be nil")
	ErrEmptyFirstNameAdministrator     = errors.New("first name cannot be empty")
	ErrEmptyLastNameAdministrator      = errors.New("last name cannot be empty")
	ErrLongFirstNameAdministrator      = errors.New("first name cannot be longer than 100 characters")
	ErrLongLastNameAdministrator       = errors.New("last name cannot be longer than 100 characters")
	ErrNonAlphaFirstNameAdministrator  = errors.New("first name has non alphabetical characters")
	ErrNonAlphaLastNameAdministrator   = errors.New("last name has non alphabetical characters")
	ErrExistAdministrator              = errors.New("administrator already exist")
	ErrNotFoundAdministrator           = errors.New("administrator not found")
	ErrInvalidCredentialsAdministrator = errors.New("invalid credentials")
)

func (a *Administrator) Id() uuid.UUID {
	return a.Entity.Id
}

func (a *Administrator) FirstName() string {
	return a.firstName
}

func (a *Administrator) LastName() string {
	return a.lastName
}

func (a *Administrator) Email() vo.Email {
	return a.email
}

func (a *Administrator) Password() vo.Password {
	return a.password
}

func (a *Administrator) Gender() vo.Gender {
	return a.gender
}

func (a *Administrator) Birth() vo.BirthDate {
	return a.birth
}

func (a *Administrator) Phone() *vo.Phone {
	return a.phone
}

func (a *Administrator) LastLoginAt() time.Time {
	return a.lastLoginAt
}

func (a *Administrator) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Administrator) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Administrator) DeletedAt() *time.Time {
	return a.deletedAt
}

func (a *Administrator) Logged() {
	a.lastLoginAt = time.Now()
}

func NewAdministrator(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone *vo.Phone) *Administrator {
	return &Administrator{
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

func NewAdministratorFromDB(id uuid.UUID, firstName string, lastName string, email string, password string, gender string, birth time.Time, phone *string, lastLoginAt time.Time, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) (*Administrator, error) {
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

	return &Administrator{
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
