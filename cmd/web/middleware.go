package main

import (
	"fmt"
	"net/http"
)

// Middleware that adds to every request headers responsible for preventing ftron XSS and Clickjack attacks.
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		
		next.ServeHTTP(w, r)
	})
}

// Middleware that logs to the console every request.
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

// Middleware that sends response user with server error instead of empty reply in case when application panic.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		defer func () {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
