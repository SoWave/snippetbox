package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// From is custom struct, which embeds a url.Values object to hold the form data
// and an Errors field to hold any validation errors from the form data.
type Form struct {
	url.Values
	Errors errors
}

// New initializes custom Form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		make(errors),
	}
}

// Required checks that specific field in the form data are present and not blank. 
// If any fields fail this check, add the appropriate message to the form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength checks that specific field in the form contains a maximum number of characters.
// If check fails, add the appropriate message to the form errors.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d)", d))
	}
}

// PermittedValues checks that specific field in the form matches one of the specific options.
// If check fails, add the appropriate message to the errors form. 
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid checks that form has no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
