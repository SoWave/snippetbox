package models

import (
	"errors"
	"time"
)

// ErrNoRecord is error with message about absence of record
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is model of db table snippets
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
