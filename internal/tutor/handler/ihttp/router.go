package ihttp

import (
	v1 "multiApp/internal/tutor/handler/ihttp/v1"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	v1.Register(r)
	return r
}
