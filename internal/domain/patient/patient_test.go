package patients

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPatient(t *testing.T) {
	phones := []string{"77141516", "77141517", "77141518", "77141519", "77141520"}
	deletedAts := []time.Time{time.Now(), time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, -2, 0), time.Now().AddDate(0, -3, 0), time.Now().AddDate(0, -4, 0)}
	cases := []struct {
		id                                                 uuid.UUID
		name, firstName, lastName, email, password, gender string
		birth, lastLoginAt, createdAt, updatedAt           time.Time
		phone                                              *string
		deletedAt                                          *time.Time
	}{
		{uuid.New(), "Case 1", "Carlos", "Clavijo", "carlos@example.com", "Abcdef1!", "M", time.Now().AddDate(-25, 0, 0), time.Now(), time.Now(), time.Now(), &phones[0], nil},
		{uuid.New(), "Case 2", "Ana", "Gomez", "ana@example.com", "Password1!", "F", time.Now().AddDate(-30, 0, 0), time.Now(), time.Now(), time.Now(), nil, &deletedAts[0]},
		{uuid.New(), "Case 3", "Luis", "Fernandez", "luis@example.com", "Strong1@", "M", time.Now().AddDate(-22, 0, 0), time.Now(), time.Now(), time.Now(), &phones[1], &deletedAts[1]},
		{uuid.New(), "Case 4", "Maria", "Lopez", "maria@example.com", "Secure2#", "F", time.Now().AddDate(-28, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
		{uuid.New(), "Case 5", "Juan", "Perez", "juan@example.com", "MyPass3$", "M", time.Now().AddDate(-35, 0, 0), time.Now(), time.Now(), time.Now(), &phones[2], &deletedAts[2]},
		{uuid.New(), "Case 6", "Lucia", "Martinez", "lucia@example.com", "TopPass4%", "F", time.Now().AddDate(-40, 0, 0), time.Now(), time.Now(), time.Now(), nil, &deletedAts[3]},
		{uuid.New(), "Case 7", "Diego", "Ramirez", "diego@example.com", "Great5^S", "M", time.Now().AddDate(-23, 0, 0), time.Now(), time.Now(), time.Now(), &phones[3], nil},
		{uuid.New(), "Case 8", "Sofia", "Torres", "sofia@example.com", "BestPass6&", "F", time.Now().AddDate(-27, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
		{uuid.New(), "Case 9", "Pedro", "Vargas", "pedro@example.com", "Power77*", "M", time.Now().AddDate(-31, 0, 0), time.Now(), time.Now(), time.Now(), &phones[4], &deletedAts[4]},
		{uuid.New(), "Case 10", "Elena", "Rojas", "elena@example.com", "Ultra8(s", "F", time.Now().AddDate(-26, 0, 0), time.Now(), time.Now(), time.Now(), nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			patient := NewPatientFromDB(tc.id, tc.firstName, tc.lastName, tc.email, tc.password, tc.gender, tc.birth, tc.phone, tc.lastLoginAt, tc.createdAt, tc.updatedAt, tc.deletedAt)

			assert.NotNil(t, patient)

			gender, _ := valueobjects.ParseGender(tc.gender)
			assert.Equal(t, tc.id, patient.Id())
			assert.Equal(t, tc.firstName, patient.FirstName())
			assert.Equal(t, tc.lastName, patient.LastName())
			assert.Equal(t, tc.email, patient.Email().Value())
			assert.Equal(t, tc.password, patient.Password().String())
			assert.Equal(t, gender.String(), patient.Gender().String())
			assert.Equal(t, tc.birth.Format("2006-01-02"), patient.Birth().Value().Format("2006-01-02"))

			if tc.phone != nil {
				assert.NotNil(t, patient.Phone())
				assert.Equal(t, *tc.phone, *patient.Phone().String())
			} else {
				assert.Nil(t, patient.Phone())
			}

			assert.Equal(t, tc.lastLoginAt.Format(time.RFC3339), patient.LastLoginAt.Format(time.RFC3339))
			assert.Equal(t, tc.createdAt.Format(time.RFC3339), patient.CreatedAt().Format(time.RFC3339))
			assert.Equal(t, tc.updatedAt.Format(time.RFC3339), patient.UpdatedAt.Format(time.RFC3339))

			if tc.deletedAt != nil {
				assert.NotNil(t, patient.DeletedAt)
				assert.Equal(t, tc.deletedAt.Format(time.RFC3339), patient.DeletedAt.Format(time.RFC3339))
			} else {
				assert.Nil(t, patient.DeletedAt)
			}
		})
	}
}
