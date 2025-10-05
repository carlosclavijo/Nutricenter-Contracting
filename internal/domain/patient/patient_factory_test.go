package patients

import (
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPatientFactory_Valid(t *testing.T) {
	factory := NewPatientFactory()

	assert.Empty(t, factory)
	assert.NotNil(t, factory)

	_, ok := factory.(PatientFactory)
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
		{"Case 10", "Matt", "Bommer", "patient@mailserver1.tv", "Go!Dev123A", "U", time.Now().AddDate(-50, 0, 0), nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := isAlpha(tc.firstName)
			lName := isAlpha(tc.lastName)

			email, err := valueobjects.NewEmail(tc.email)
			assert.NotEmpty(t, email)
			assert.Equal(t, email.Value(), tc.email)
			assert.Nil(t, err)

			password, er := valueobjects.NewPassword(tc.password)
			assert.NotEmpty(t, password)
			assert.Equal(t, password.String(), tc.password)
			assert.Nil(t, er)

			gender, err := valueobjects.ParseGender(tc.gender)
			assert.Contains(t, []string{"undefined", "male", "female"}, gender.String())
			assert.NotEqual(t, gender, "")
			assert.NotEqual(t, gender, "unknown")
			assert.NoError(t, err)

			birth, err := valueobjects.NewBirthDate(tc.birth)
			assert.NotEmpty(t, birth)
			assert.Equal(t, birth.Value(), tc.birth)
			assert.Nil(t, err)

			phone, err := valueobjects.NewPhone(tc.phone)
			if phone == nil {
				assert.Nil(t, err)
			} else {
				assert.NotEmpty(t, phone)
				assert.Equal(t, phone.String(), tc.phone)
				assert.Nil(t, err)
			}

			patient, err := factory.Create(tc.firstName, tc.lastName, email, password, gender, birth, phone)

			assert.NotNil(t, patient)
			assert.Nil(t, err)
			assert.NoError(t, err)

			assert.True(t, fName)
			assert.True(t, lName)

			assert.Equal(t, tc.firstName, patient.FirstName())
			assert.Equal(t, tc.lastName, patient.LastName())
			assert.Equal(t, tc.email, patient.Email().Value())
			assert.Equal(t, tc.password, patient.Password().String())
			assert.Equal(t, gender.String(), patient.Gender().String())
			assert.Contains(t, []string{"undefined", "male", "female"}, patient.Gender().String())
			assert.Equal(t, tc.birth.Format("2006-01-02"), patient.Birth().Value().Format("2006-01-02"))

			if tc.phone != nil {
				assert.NotNil(t, patient.Phone())
				assert.Equal(t, *tc.phone, *patient.Phone().String())
			} else {
				assert.Nil(t, patient.Phone())
			}

			assert.NotNil(t, patient.Id())
			assert.NotNil(t, patient.LastLoginAt)
			assert.Empty(t, patient.LastLoginAt)
			assert.NotNil(t, patient.CreatedAt())
			assert.Empty(t, patient.CreatedAt())
			assert.NotNil(t, patient.UpdatedAt)
			assert.NotNil(t, patient.UpdatedAt)
			assert.Nil(t, patient.DeletedAt)
		})
	}
}

func TestPatientFactory_Invalid_Empty(t *testing.T) {
	factory := NewPatientFactory()
	patient, err := factory.Create("", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "firstName is empty")

	patient, err = factory.Create("Carlos", "", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "lastName is empty")
}

func TestPatientFactory_Invalid_LongNames(t *testing.T) {
	factory := NewPatientFactory()
	name := "ThisNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	patient, err := factory.Create(name, "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)

	name = "ThisLastNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	patient, err = factory.Create("Carlos", name, valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
	assert.NotNil(t, err)

	expected = fmt.Sprintf("lastName '%s' is too long('%d'), maximum length is 100 characters", name, len(name))
	assert.ErrorContains(t, err, expected)
}

func TestPatientFactory_Invalid_NonAlpha(t *testing.T) {
	factory := NewPatientFactory()
	patient, err := factory.Create("Carlos123", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
	assert.NotNil(t, err)

	expected := fmt.Sprintf("firstName '%s' contains non-alphabetic characters", "Carlos123")
	assert.ErrorContains(t, err, expected)

	patient, err = factory.Create("Carlos", "Clavijo!", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)

	assert.Nil(t, patient)
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
