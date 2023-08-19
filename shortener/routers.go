package shortener

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewMux()
	templates := template.Must(template.ParseGlob("template/*.html"))
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Execute(w, nil)
	})

	r.Post("/api/create", createLink)
	r.Get("/{url}", redirectLink)
	return r
}
