package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handle(r *chi.Mux) *chi.Mux {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
	})

	return r
}
