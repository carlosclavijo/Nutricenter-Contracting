package administrators

import (
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
	gender      string
	birth       *vo.BirthDate
	phone       *vo.Phone
	lastLoginAt *time.Time
	createdAt   *time.Time
	updatedAt   *time.Time
	deletedAt   *time.Time
}

func NewAdministrator(firstName, lastName string, email vo.Email, password vo.Password, gender string, birth *vo.BirthDate, phone *vo.Phone) *Administrator {
	return &Administrator{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		password:      password,
		gender:        gender,
		birth:         birth,
		phone:         phone,
		lastLoginAt:   nil,
		createdAt:     nil,
		updatedAt:     nil,
		deletedAt:     nil,
	}
}

func (admin *Administrator) Id() uuid.UUID {
	return admin.Entity.Id
}

func (admin *Administrator) FirstName() string {
	return admin.firstName
}

func (admin *Administrator) LastName() string {
	return admin.lastName
}

func (admin *Administrator) Email() vo.Email {
	return admin.email
}

func (admin *Administrator) Password() vo.Password {
	return admin.password
}

func (admin *Administrator) Gender() string {
	return admin.gender
}

func (admin *Administrator) Birth() *vo.BirthDate {
	return admin.birth
}

func (admin *Administrator) Phone() *vo.Phone {
	return admin.phone
}

func (admin *Administrator) LastLoginAt() *time.Time {
	return admin.lastLoginAt
}

func (admin *Administrator) CreatedAt() *time.Time {
	return admin.createdAt
}

func (admin *Administrator) UpdatedAt() *time.Time {
	return admin.updatedAt
}

func (admin *Administrator) DeletedAt() *time.Time {
	return admin.deletedAt
}

func NewAdministratorFromDB(id uuid.UUID, firstName, lastName, email, password, gender string, birth *time.Time, phone *string, lastLoginAt, createdAt, updatedAt, deletedAt *time.Time) *Administrator {
	emailVo, _ := vo.NewEmail(email)
	passwordVo, _ := vo.NewPassword(password)
	birthVo, _ := vo.NewBirthDate(birth)
	phoneVo, _ := vo.NewPhone(phone)

	return &Administrator{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		firstName:     firstName,
		lastName:      lastName,
		email:         emailVo,
		password:      passwordVo,
		gender:        gender,
		birth:         birthVo,
		phone:         phoneVo,
		lastLoginAt:   lastLoginAt,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}
}
