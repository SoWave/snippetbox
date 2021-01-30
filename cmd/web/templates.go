package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/SoWave/snippetbox/pkg/forms"
	"github.com/SoWave/snippetbox/pkg/models"
)

// Contains data used in templates.
type templateData struct {
	CurrentYear int
	Flash       string
	Form        *forms.Form
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// HumanDate formats time input to readable form.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Passes functions to templates.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Returns template cache containing assembled templates.
// Cache is helpful when we don't need new request every time we change template.
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
