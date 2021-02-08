package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("Want body to equal %q", "OK")
	}
}

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	testCases := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old silent pond...")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			code, _, body := ts.get(t, tC.urlPath)

			if code != tC.wantCode {
				t.Errorf("want %d; got %d", tC.wantCode, code)
			}

			if !bytes.Contains(body, tC.wantBody) {
				t.Errorf("want body to contain %q", tC.wantBody)
			}
		})
	}
}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	testCases := []struct {
		desc         string
		userName     string
		userEmail    string
		userPassword string
		CSRFToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", "Bob", "bob@example.com", "validPa$$word", csrfToken, http.StatusSeeOther, []byte("")},
		{"Empty name", "", "bob@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty email", "Bob", "", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty password", "Bob", "bob@example.com", "", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Invalid email (incomplete domain)", "Bob", "bob@example.", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", "Bob", "bobexample.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing local part)", "Bob", "@example.com", "validPa$$", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", "Bob", "bob@example.com", "pa$$word", csrfToken, http.StatusOK, []byte("This field is too short (minimum is 10)")},
		{"Duplicate email", "Bob", "dpl@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, []byte("")},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tC.userName)
			form.Add("email", tC.userEmail)
			form.Add("password", tC.userPassword)
			form.Add("csrf_token", tC.CSRFToken)

			code, _, body := ts.post(t, "/user/signup", form)

			if code != tC.wantCode {
				t.Errorf("want %d; got %d", tC.wantCode, code)
			}

			if !bytes.Contains(body, tC.wantBody) {
				t.Errorf("want body %s to contain %q", body, tC.wantBody)
			}
		})
	}
}
