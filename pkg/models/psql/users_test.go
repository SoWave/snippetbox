package psql

import (
	"testing"
	"time"

	"github.com/SoWave/snippetbox/pkg/models"
)

func TestUserModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	testCases := []struct {
		desc      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			desc:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Alice Jones",
				Email:   "alice@example.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			desc:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			desc:      "Non-existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := UserModel{db}
			user, err := m.Get(tC.userID)

			if err != tC.wantError {
				t.Errorf("want %v; got %s", tC.wantError, err)
			}

			if user != nil && (user.ID != tC.wantUser.ID || !user.Created.Equal(tC.wantUser.Created) ||
				user.Email != tC.wantUser.Email || user.Name != tC.wantUser.Name) {

				t.Errorf("want %v; got %v", tC.wantUser, user)
			}
		})
	}
}
