package http

import (
	"net/http"
	"proj/internal/domain"
	ctxkey "proj/pkg/context_key"
	resp "proj/pkg/responses"

	"github.com/go-chi/render"
)

type NotesService interface {
	CreateNote(text string, userID string) (*domain.Note, error)
	GetNote(userID string, noteID string) (*domain.Note, error)
	UpdateNote(userID string, noteID string, text string) error
	DeleteNote(userID string, noteID string) error
	ListUserNotes(userID string) ([]*domain.Note, error)
}

type NotesHandler struct {
	service NotesService
}

func NewNotesHandler(service NotesService) *NotesHandler {
	return &NotesHandler{service: service}
}

type createNoteRequest struct {
	Text string `json:"text"`
}

func (h *NotesHandler) createNote(w http.ResponseWriter, r *http.Request) {
	in := &createNoteRequest{}
	if err := render.Decode(r, in); err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	userID := ctxkey.GetFormCtx(r, userIDCtxKey)
	note, err := h.service.CreateNote(in.Text, userID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, note)
}

func (h *NotesHandler) getNote(w http.ResponseWriter, r *http.Request) {
	userID := ctxkey.GetFormCtx(r, userIDCtxKey)
	noteID := ctxkey.GetFormCtx(r, noteIDCtxKey)
	note, err := h.service.GetNote(userID, noteID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, note)
}

type updateNoteRequest struct {
	Text string `json:"text"`
}

func (h *NotesHandler) updateNote(w http.ResponseWriter, r *http.Request) {
	in := &updateNoteRequest{}
	if err := render.Decode(r, in); err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	userID := ctxkey.GetFormCtx(r, userIDCtxKey)
	noteID := ctxkey.GetFormCtx(r, noteIDCtxKey)
	err := h.service.UpdateNote(userID, noteID, in.Text)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, resp.OK())
}

func (h *NotesHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	userID := ctxkey.GetFormCtx(r, userIDCtxKey)
	noteID := ctxkey.GetFormCtx(r, noteIDCtxKey)
	err := h.service.DeleteNote(userID, noteID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, resp.OK())
}

type listNotesResponse struct {
	Count int            `json:"count"`
	Notes []*domain.Note `json:"notes"`
}

func (h *NotesHandler) listNotes(w http.ResponseWriter, r *http.Request) {
	userID := ctxkey.GetFormCtx(r, userIDCtxKey)
	userNotes, err := h.service.ListUserNotes(userID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, listNotesResponse{Count: len(userNotes), Notes: userNotes})
}
