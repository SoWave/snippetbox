package mock

import (
	"time"

	"github.com/SoWave/snippetbox/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel mock.
type SnippetModel struct{}

// Insert mock impl.
func (m *SnippetModel) Insert(title, content string, expires []byte) (int, error) {
	return 2, nil
}

// Get mock impl.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest mock impl
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
