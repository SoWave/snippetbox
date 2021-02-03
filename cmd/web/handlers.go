package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SoWave/snippetbox/pkg/forms"
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

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// CreateSnippetForm handler renders form for creating new snippet
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// Create snippet handler adds snippet from form.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	expires, err := strconv.Atoi(form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	fExpires := pq.FormatTimestamp(time.Now().AddDate(0, 0, expires))

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), fExpires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display user signup form..")
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create new user...")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display user login form...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
