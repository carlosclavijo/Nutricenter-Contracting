package commands

import "time"

type CreateAdministratorCommand struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    string
	Birth     time.Time
	Phone     *string
}
