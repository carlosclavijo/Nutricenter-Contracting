package patients

import (
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
	gender      string
	birth       *vo.BirthDate
	phone       *vo.Phone
	LastLoginAt time.Time
	createdAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewPatient(firstName, lastName string, email vo.Email, password vo.Password, gender string, birth *vo.BirthDate, phone *vo.Phone) *Patient {
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

func (admin *Patient) Id() uuid.UUID {
	return admin.Entity.Id
}

func (admin *Patient) FirstName() string {
	return admin.firstName
}

func (admin *Patient) LastName() string {
	return admin.lastName
}

func (admin *Patient) Email() vo.Email {
	return admin.email
}

func (admin *Patient) Password() vo.Password {
	return admin.password
}

func (admin *Patient) Gender() string {
	return admin.gender
}

func (admin *Patient) Birth() *vo.BirthDate {
	return admin.birth
}

func (admin *Patient) Phone() *vo.Phone {
	return admin.phone
}

func (admin *Patient) CreatedAt() time.Time {
	return admin.createdAt
}

func NewPatientFromDB(id uuid.UUID, firstName, lastName, email, password, gender string, birth *time.Time, phone *string, lastLoginAt, createdAt, updatedAt time.Time, deletedAt *time.Time) *Patient {
	emailVo, _ := vo.NewEmail(email)
	passwordVo, _ := vo.NewPassword(password)
	birthVo, _ := vo.NewBirthDate(birth)
	phoneVo, _ := vo.NewPhone(phone)

	return &Patient{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		firstName:     firstName,
		lastName:      lastName,
		email:         emailVo,
		password:      passwordVo,
		gender:        gender,
		birth:         birthVo,
		phone:         phoneVo,
		LastLoginAt:   lastLoginAt,
		createdAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}
