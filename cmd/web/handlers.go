package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SoWave/snippetbox/pkg/models"
	
	"github.com/lib/pq"
)

// Home handler "/" renders home page with 10 latest snippets.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

// Show snippet handler render show page with info about snippet with {?=id}.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

// CreateSnippetForm handler renders form with new snippet data
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet"))
}

// Create snippet handler adds snippet with given data.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "O Snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := time.Now().AddDate(0, 0, 7)
	formatedExp := pq.FormatTimestamp(expires)

	id, err := app.snippets.Insert(title, content, formatedExp)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
