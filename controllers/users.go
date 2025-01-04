package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email") // Try to get email from URL query params (could be handy if user was forwarded from other site)
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "email: %v", r.FormValue("email"))
	fmt.Fprintf(w, "password: %v", r.FormValue("password"))
}
