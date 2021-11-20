package http

import (
	"proj/internal/service"
	ctxKey "proj/pkg/context_key"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	services service.Services
}

func NewRouter(s service.Services) *Router {
	return &Router{
		services: s,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/users", r.initUserRouter())
	return router
}

func (r *Router) initUserRouter() chi.Router {
	handler := NewUsersHandler(r.services.UserService)
	router := chi.NewRouter()
	router.Post("/", handler.createUser)
	router.Route("/{userID}", func(subRouter chi.Router) {
		subRouter.Use(ctxKey.URLParamCtx("userID", userIDCtxKey))
		subRouter.Get("/", handler.getUser)
		subRouter.Delete("/", handler.deleteUser)
		subRouter.Mount("/notes", r.initNotesRouter())
	})

	return router
}

func (r *Router) initNotesRouter() chi.Router {
	handler := NewNotesHandler(r.services.NotesService)
	router := chi.NewRouter()
	router.Get("/", handler.listNotes)
	router.Post("/", handler.createNote)
	router.Route("/{noteID}", func(subRouter chi.Router) {
		subRouter.Use(ctxKey.URLParamCtx("noteID", noteIDCtxKey))
		subRouter.Get("/", handler.getNote)
		subRouter.Patch("/", handler.updateNote)
		subRouter.Delete("/", handler.deleteNote)
	})
	return router
}
