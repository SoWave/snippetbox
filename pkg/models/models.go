package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord error with message about absence of record.
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCreditentials error with message about invalid creditantials.
	ErrInvalidCreditentials = errors.New("models: invalid creditentials")

	// ErrDuplicateEmail error with message about duplicate email.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet model in snippets db
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User model in snippets db
type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
	Created  time.Time
}
