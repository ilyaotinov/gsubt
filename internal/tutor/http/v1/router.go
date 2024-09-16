package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handle(r *chi.Mux) *chi.Mux {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
	})

	return r
}
