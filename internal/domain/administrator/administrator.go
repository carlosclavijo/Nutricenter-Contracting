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
	gender      vo.Gender
	birth       vo.BirthDate
	phone       *vo.Phone
	lastLoginAt time.Time
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
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

func NewAdministratorFromDB(id uuid.UUID, firstName string, lastName string, email string, password string, gender string, birth time.Time, phone *string, lastLoginAt time.Time, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *Administrator {
	emailVo, _ := vo.NewEmail(email)
	passwordVo, _ := vo.NewPassword(password)
	genderVo, _ := vo.ParseGender(gender)
	birthVo, _ := vo.NewBirthDate(birth)
	phoneVo, _ := vo.NewPhone(phone)

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
	}
}
