package main

import (
	"net/url"
	"strings"
)

// errors is a convenience type, sho that we can have a function tied to our map
type errors map[string][]string

func (e errors) Get(field string) string {
	errorSlice := e[field]

	if len(errorSlice) == 0 {
		return ""
	}
	return errorSlice[0]
}

//add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//form is the type used to instantie form validation
type Form struct {
	Data   url.Values
	Errors errors
}

//newForm initializes a form struct
func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

//has check to see if the form has a given fields
func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	if x == "" {
		return false
	}
	return true
}

//required checks  for require fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//check is a generic validation checks, we can pass expression
//that evaluates aas a boolean as the first parameter
func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.Errors.Add(key, message)
	}
}

// vaid returns true if there are no errors, otherwis false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
