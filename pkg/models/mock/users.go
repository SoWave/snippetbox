package mock

import (
	"time"

	"github.com/SoWave/snippetbox/pkg/models")


var mockUser = &models.User{
	ID: 1,
	Name: "Alice",
	Email: "alice@example.com",
	Created: time.Now(),
}

// UserModel mock.
type UserModel struct{}

// Insert mock impl.
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dpl@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate mock impl.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@examle.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCreditentials
	}
}

// Get mock impl.
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
