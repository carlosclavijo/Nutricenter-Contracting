package administrators

import (
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdministratorFactory(t *testing.T) {
	factory := NewAdministratorFactory()

	assert.NotNil(t, factory)
	assert.Empty(t, factory)

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

			assert.True(t, fName)
			assert.True(t, lName)

			email, err := valueobjects.NewEmail(tc.email)
			assert.NotEmpty(t, email)
			assert.Equal(t, tc.email, email.Value())
			assert.Nil(t, err)

			password, err := valueobjects.NewPassword(tc.password)
			assert.NotEmpty(t, password)
			assert.Equal(t, tc.password, password.String())
			assert.Nil(t, err)

			gender, err := valueobjects.ParseGender(tc.gender)
			assert.Contains(t, []string{"undefined", "male", "female"}, gender.String())
			assert.NotEqual(t, "", gender)
			assert.NotEqual(t, "unknown", gender)
			assert.NoError(t, err)

			birth, err := valueobjects.NewBirthDate(tc.birth)
			assert.NotEmpty(t, birth)
			assert.Equal(t, tc.birth, birth.Value())
			assert.Nil(t, err)

			phone, err := valueobjects.NewPhone(tc.phone)
			if phone == nil {
				assert.Nil(t, err)
			} else {
				assert.NotEmpty(t, phone)
				assert.Equal(t, tc.phone, phone.String())
				assert.Nil(t, err)
			}

			admin, err := factory.Create(tc.firstName, tc.lastName, email, password, gender, birth, phone)
			assert.NotNil(t, admin)
			assert.NotNil(t, admin.Id())
			assert.NotNil(t, admin.LastLoginAt())
			assert.NotNil(t, admin.CreatedAt())
			assert.NotNil(t, admin.UpdatedAt())

			assert.Equal(t, tc.firstName, admin.FirstName())
			assert.Equal(t, tc.lastName, admin.LastName())
			assert.Equal(t, tc.email, admin.Email().Value())
			assert.Equal(t, tc.password, admin.Password().String())
			assert.Equal(t, gender.String(), admin.Gender().String())
			assert.Contains(t, []string{"undefined", "male", "female"}, admin.Gender().String())
			assert.Equal(t, tc.birth.Format(time.RFC3339), admin.Birth().Value().Format(time.RFC3339))

			if tc.phone != nil {
				assert.NotNil(t, admin.Phone())
				assert.Equal(t, *tc.phone, *admin.Phone().String())
			} else {
				assert.Nil(t, admin.Phone())
			}

			assert.Empty(t, admin.LastLoginAt())
			assert.Empty(t, admin.CreatedAt())
			assert.Empty(t, admin.UpdatedAt())
			assert.Nil(t, admin.DeletedAt())

			assert.Nil(t, err)
			assert.NoError(t, err)
		})
	}
}

func TestAdministratorFactory_EmptyError(t *testing.T) {
	factory := NewAdministratorFactory()
	admin, err := factory.Create("", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "firstName is empty")
	assert.Nil(t, admin)

	admin, err = factory.Create("Carlos", "", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "lastName is empty")
	assert.Nil(t, admin)
}

func TestAdministratorFactory_LongNameError(t *testing.T) {
	factory := NewAdministratorFactory()
	name := "ThisNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	admin, err := factory.Create(name, "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)

	assert.Nil(t, admin)

	name = "ThisLastNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	admin, err = factory.Create("Carlos", name, valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.NotNil(t, err)

	expected = fmt.Sprintf("lastName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)

	assert.Nil(t, admin)
}

func TestAdministratorFactory_NonAlphaEror(t *testing.T) {
	factory := NewAdministratorFactory()
	admin, err := factory.Create("Carlos123", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' contains non-alphabetic characters", "Carlos123")
	assert.ErrorContains(t, err, expected)

	assert.Nil(t, admin)

	admin, err = factory.Create("Carlos", "Clavijo!", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.NotNil(t, err)

	expected = fmt.Sprintf("lastName '%s' contains non-alphabetic characters", "Clavijo!")
	assert.ErrorContains(t, err, expected)

	assert.Nil(t, admin)
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
