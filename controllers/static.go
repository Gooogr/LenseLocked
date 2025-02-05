package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	// because of closure we keep access to tpl object in StaticHanlder level
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
