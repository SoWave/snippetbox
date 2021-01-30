package forms

// New errors type. Field of the form is used as key.
type errors map[string][]string

// Add function describes error by field with message then ad to the errors stack.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get retrieves first error message from given field.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}