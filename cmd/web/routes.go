package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// Routes defines routes to handlers and static server
func (app *application) routes() http.Handler {
	// Middleware
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamic := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamic.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamic.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamic.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamic.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamic.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamic.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamic.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamic.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamic.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standard.Then(mux)
}