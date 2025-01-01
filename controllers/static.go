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

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Question 1",
			Answer:   "Answer 1",
		},
		{
			Question: "Question 2",
			Answer:   "Answer 2",
		},
		{
			Question: "Question 3",
			Answer:   "Answer 3",
		},
		{
			Question: "Question 4",
			Answer:   "Answer 4",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
