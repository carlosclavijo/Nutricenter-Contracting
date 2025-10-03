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
	LastLoginAt time.Time
	createdAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
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

func (admin *Administrator) Gender() vo.Gender {
	return admin.gender
}

func (admin *Administrator) Birth() vo.BirthDate {
	return admin.birth
}

func (admin *Administrator) Phone() *vo.Phone {
	return admin.phone
}

func (admin *Administrator) CreatedAt() time.Time {
	return admin.createdAt
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
		LastLoginAt:   lastLoginAt,
		createdAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}
