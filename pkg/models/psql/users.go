package psql

import (
	"database/sql"

	"github.com/SoWave/snippetbox/pkg/models"
)

// UserModel wraps sql.DB connection pool
type UserModel struct {
	DB *sql.DB
}

// Insert new user to the database.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate that user with specified email and password exists.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get specific user based on their id.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}