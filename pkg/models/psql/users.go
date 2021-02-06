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
	var id int
	var hashedPass []byte

	err := m.DB.QueryRow("SELECT id, password FROM USERS WHERE email = $1", email).Scan(&id, &hashedPass)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCreditentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCreditentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// Get specific user based on their id.
func (m *UserModel) Get(id int) (*models.User, error) {
	s := &models.User{}

	stmt := "SELECT id, name, email, created FROM users WHERE id = $1"
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.Email, &s.Created)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
