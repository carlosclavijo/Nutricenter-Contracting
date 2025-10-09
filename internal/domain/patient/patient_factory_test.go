package patients

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPatientFactory(t *testing.T) {
	factory := NewPatientFactory()

	assert.NotNil(t, factory)
	assert.Empty(t, factory)

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

			assert.True(t, fName)
			assert.True(t, lName)

			email, err := valueobjects.NewEmail(tc.email)
			assert.NotEmpty(t, email)
			assert.Equal(t, tc.email, email.Value())
			assert.Nil(t, err)

			password, er := valueobjects.NewPassword(tc.password)
			assert.NotEmpty(t, password)
			assert.Equal(t, tc.password, password.String())
			assert.Nil(t, er)

			gender, err := valueobjects.ParseGender(tc.gender)
			assert.Contains(t, []string{"undefined", "male", "female"}, gender.String())
			assert.NotEqual(t, gender, "")
			assert.NotEqual(t, gender, "unknown")
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

			patient, err := factory.Create(tc.firstName, tc.lastName, email, password, gender, birth, phone)

			assert.NotNil(t, patient)
			assert.NotNil(t, patient.Id())
			assert.NotNil(t, patient.LastLoginAt())
			assert.NotNil(t, patient.CreatedAt())
			assert.NotNil(t, patient.UpdatedAt())

			assert.Equal(t, tc.firstName, patient.FirstName())
			assert.Equal(t, tc.lastName, patient.LastName())
			assert.Equal(t, tc.email, patient.Email().Value())
			assert.Equal(t, tc.password, patient.Password().String())
			assert.Equal(t, gender.String(), patient.Gender().String())
			assert.Contains(t, []string{"undefined", "male", "female"}, patient.Gender().String())
			assert.Equal(t, patient.Birth().Value().Format(time.RFC3339), tc.birth.Format(time.RFC3339))

			if tc.phone != nil {
				assert.NotNil(t, patient.Phone())
				assert.Equal(t, *patient.Phone().String(), *tc.phone)
			} else {
				assert.Nil(t, patient.Phone())
			}

			assert.Empty(t, patient.LastLoginAt())
			assert.Empty(t, patient.CreatedAt())
			assert.Empty(t, patient.UpdatedAt())
			assert.Nil(t, patient.DeletedAt())
			assert.Nil(t, err)
			assert.NoError(t, err)
		})
	}
}

func TestPatientFactory_EmptyError(t *testing.T) {
	factory := NewPatientFactory()
	patient, err := factory.Create("", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrEmptyFirstNamePatient)
	assert.Nil(t, patient)

	patient, err = factory.Create("Carlos", "", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrEmptyLastNamePatient)
	assert.Nil(t, patient)
}

func TestPatientFactory_LongNameError(t *testing.T) {
	factory := NewPatientFactory()
	name := "ThisNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	patient, err := factory.Create(name, "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrLongFirstNamePatient)
	assert.Nil(t, patient)

	name = "ThisLastNameIsWayTooLongToBeConsideredValidBecauseItExceedsTheMaximumLengthOfOneHundredCharactersWhichIsNotAllowed"
	patient, err = factory.Create("Carlos", name, valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrLongLastNamePatient)
	assert.Nil(t, patient)
}

func TestPatientFactory_NonAlphaError(t *testing.T) {
	factory := NewPatientFactory()
	patient, err := factory.Create("Carlos123", "Clavijo", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrNonAlphaFirstNamePatient)
	assert.Nil(t, patient)

	patient, err = factory.Create("Carlos", "Clavijo!", valueobjects.Email{}, valueobjects.Password{}, "", valueobjects.BirthDate{}, nil)
	assert.ErrorIs(t, err, ErrNonAlphaLastNamePatient)
	assert.Nil(t, patient)
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
