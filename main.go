package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gooogr/lenselocked/controllers"
	"github.com/gooogr/lenselocked/views"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", r)
}
