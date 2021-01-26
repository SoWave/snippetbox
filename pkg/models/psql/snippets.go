package psql

import (
	"database/sql"

	"github.com/SoWave/snippetbox/pkg/models"
)

// SnippetModel type wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet to database.
func (m *SnippetModel) Insert(title, content string, expires []byte) (int, error) {
	stmt := `INSERT INTO snippets (title, snippet_content, created, expires)
	VALUES($1, $2, NOW(), $3)
	RETURNING snippet_id`

	var snippetid int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&snippetid)
	if err != nil {
		return 0, err
	}

	return snippetid, nil
}

// Get snippet with id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT * FROM snippets WHERE expires > NOW() AND snippet_id=$1`

	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord	
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Latest return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT snippet_id, title, snippet_content, created, expires FROM snippets
	WHERE expires > NOW() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
