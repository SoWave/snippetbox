package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"

	"github.com/SoWave/snippetbox/pkg/models/psql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "web"
	password = "pass"
	dbname   = "snippetbox"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *psql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	fmt.Println(templateCache)

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &psql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(info string) (*sql.DB, error) {
	db, err := sql.Open("postgres", info)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
