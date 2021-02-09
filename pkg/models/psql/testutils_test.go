package psql

import (
	"database/sql"
	"io/ioutil"
	"testing"
)

// NewTestDB creates new connection pool to test database. Creates tables from setup.sql, inserts dummy data
// and returns db connection with teardown function that drops every created data and close the connection.
func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("postgres", `host=localhost port=5432 user=test_web 
		password=test dbname=test_snippetbox sslmode=disable`)
	if err != nil {
		t.Fatal(err)
	}

	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := ioutil.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	}
}
