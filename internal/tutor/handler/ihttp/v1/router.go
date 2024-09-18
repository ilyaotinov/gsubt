package v1

import (
	"github.com/go-chi/chi/v5"
)

func Register(r *chi.Mux) *chi.Mux {
	r.Route("/api/v1", func(r chi.Router) {
	})

	return r
}
