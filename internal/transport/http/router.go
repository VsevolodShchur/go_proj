package http

import (
	ctxKey "proj/pkg/context_key"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	handlers *Handlers
}

func NewRouter(handlers *Handlers) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/users", r.initUserRouter())
	return router
}

func (r *Router) initUserRouter() chi.Router {
	router := chi.NewRouter()
	router.Post("/", r.handlers.usersHandler.createUser)
	router.Route("/{userID}", func(subRouter chi.Router) {
		subRouter.Use(ctxKey.URLParamCtx("userID", userIDCtxKey))
		subRouter.Get("/", r.handlers.usersHandler.getUser)
		subRouter.Delete("/", r.handlers.usersHandler.deleteUser)
		subRouter.Mount("/notes", r.initNotesRouter())
	})

	return router
}

func (r *Router) initNotesRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/", r.handlers.notesHandler.listNotes)
	router.Post("/", r.handlers.notesHandler.createNote)
	router.Route("/{noteID}", func(subRouter chi.Router) {
		subRouter.Use(ctxKey.URLParamCtx("noteID", noteIDCtxKey))
		subRouter.Get("/", r.handlers.notesHandler.getNote)
		subRouter.Patch("/", r.handlers.notesHandler.updateNote)
		subRouter.Delete("/", r.handlers.notesHandler.deleteNote)
	})
	return router
}
