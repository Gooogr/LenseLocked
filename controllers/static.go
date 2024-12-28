package controllers

import (
	"github.com/gooogr/lenselocked/views"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
