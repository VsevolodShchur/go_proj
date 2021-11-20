package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	return r
}
