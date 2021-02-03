package psql

import (
	"database/sql"
	"strings"

	"github.com/SoWave/snippetbox/pkg/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// UserModel wraps sql.DB connection pool
type UserModel struct {
	DB *sql.DB
}

// Insert new user to the database.
func (m *UserModel) Insert(name, email, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, password, created) 
		VALUES($1, $2, $3, NOW())`
	
	_, err = m.DB.Exec(stmt, name, email, hashedPass)
	if err != nil {
		if psqlErr, ok := err.(*pq.Error); ok {
			if psqlErr.Code == "23505" && strings.Contains(psqlErr.Message, "duplicate key value violates unique constraint") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

// Authenticate that user with specified email and password exists.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get specific user based on their id.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
