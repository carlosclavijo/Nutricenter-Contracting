package administrators

import (
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdministratorFactory_Valid(t *testing.T) {
	factory := NewAdministratorFactory()

	assert.Empty(t, factory)
	assert.NotNil(t, factory)

	_, ok := factory.(AdministratorFactory)
	assert.True(t, ok)

	phones := []string{"77141516", "77141517", "77141518", "77141519", "77141520"}
	cases := []struct {
		name, firstName, lastName, email, password, gender string
		birth                                              time.Time
		phone                                              *string
	}{
		{"Case 1", "Carlos", "Clavijo", "user@example.com", "Abcdef1!", "M", time.Now().AddDate(-20, 0, 0), &phones[0]},
		{"Case 2", "Juan", "Pérez", "user.name@example.com", "GoLangR0cks@", "male", time.Now().AddDate(-25, 0, 0), &phones[1]},
		{"Case 3", "Giovanni", "Giorgio", "user_name@example.co.uk", "P@ssw0rd123", "M", time.Now().AddDate(-19, 0, 0), &phones[2]},
		{"Case 4", "Aphex", "Twin", "user-name+tag@example.io", "Str0ng#Key", "male", time.Now().AddDate(-18, -1, 0), &phones[3]},
		{"Case 5", "Marie", "Curie", "user123@example-domain.com", "Test!2024A", "female", time.Now().AddDate(-50, 0, 0), &phones[4]},
		{"Case 6", "Lana", "Del Rey", "x@example.com", "My_Secure9?", "F", time.Now().AddDate(-50, 0, 0), nil},
		{"Case 7", "Sabrina", "Carpenter", "very.common@example.com", "XyZ@98765a", "female", time.Now().AddDate(-50, 0, 0), nil},
		{"Case 8", "Taylor", "Swift", "disposable.style.email.with+symbol@example.com", "SafePass!1", "F", time.Now().AddDate(-50, 0, 0), nil},
		{"Case 9", "Eliot", "Page", "other.email-with-dash@example.com", "ValidKey#2025", "undefined", time.Now().AddDate(-50, 0, 0), nil},
		{"Case 10", "Matt", "Bommer", "admin@mailserver1.tv", "Go!Dev123A", "U", time.Now().AddDate(-50, 0, 0), nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := isAlpha(tc.firstName)
			lName := isAlpha(tc.lastName)
			email, _ := valueobjects.NewEmail(tc.email)
			password, _ := valueobjects.NewPassword(tc.password)
			gender, _ := valueobjects.ParseGender(tc.gender)
			birth, _ := valueobjects.NewBirthDate(tc.birth)
			phone, _ := valueobjects.NewPhone(tc.phone)

			admin, err := factory.Create(tc.firstName, tc.lastName, email, password, gender, birth, phone)

			assert.NotNil(t, admin)
			assert.Nil(t, err)
			assert.NoError(t, err)

			assert.True(t, fName)
			assert.True(t, lName)

			assert.Equal(t, tc.firstName, admin.FirstName())
			assert.Equal(t, tc.lastName, admin.LastName())
			assert.Equal(t, tc.email, admin.Email().Value())
			assert.Equal(t, tc.password, admin.Password().String())
			assert.Equal(t, gender.String(), admin.Gender().String())
			assert.Equal(t, tc.birth.Format("2006-01-02"), admin.Birth().Value().Format("2006-01-02"))
			assert.NotNil(t, admin.Id())
			assert.NotNil(t, admin.CreatedAt())

			if tc.phone != nil {
				assert.NotNil(t, admin.Phone())
				assert.Equal(t, *tc.phone, *admin.Phone().String())
			} else {
				assert.Nil(t, admin.Phone())
			}
		})
	}
}

func TestAdministratorFactory_Invalid_Empty(t *testing.T) {
	factory := NewAdministratorFactory()
	admin, err := factory.Create("", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "firstName is empty")

	admin, err = factory.Create("Carlos", "", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "lastName is empty")
}

func TestAdministratorFactory_Invalid_LongNames(t *testing.T) {
	factory := NewAdministratorFactory()
	name := "ThisNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	admin, err := factory.Create(name, "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)

	name = "ThisLastNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	admin, err = factory.Create("Carlos", name, valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)

	expected = fmt.Sprintf("lastName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)
}

func TestAdministratorFactory_Invalid_NonAlpha(t *testing.T) {
	factory := NewAdministratorFactory()
	admin, err := factory.Create("Carlos123", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' contains non-alphabetic characters", "Carlos123")
	assert.ErrorContains(t, err, expected)

	admin, err = factory.Create("Carlos", "Clavijo!", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, admin)
	assert.NotNil(t, err)

	expected = fmt.Sprintf("lastName '%s' contains non-alphabetic characters", "Clavijo!")
	assert.ErrorContains(t, err, expected)
}

func TestIsAlpha(t *testing.T) {
	cases := []struct {
		name, input string
		expected    bool
	}{
		{"Valid: Simple", "Carlos", true},
		{"Valid: With Space", "Carlos Alberto", true},
		{"Valid: With Multiple Spaces", "Carlos   Alberto", false},
		{"Invalid: With Number", "Carlos123", false},
		{"Invalid: With Special Char", "Carlos!", false},
		{"Invalid: Empty String", "", false},
		{"Invalid: Only Spaces", "     ", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := isAlpha(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
