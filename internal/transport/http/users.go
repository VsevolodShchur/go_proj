package http

import (
	"net/http"
	"proj/internal/domain"
	ctxKey "proj/pkg/context_key"
	resp "proj/pkg/responses"

	"github.com/go-chi/render"
)

type UserService interface {
	CreateUser(name string) (*domain.User, error)
	GetUser(ID string) (*domain.User, error)
	DeleteUser(ID string) error
}

type UsersHandler struct {
	service UserService
}

func NewUsersHandler(service UserService) *UsersHandler {
	return &UsersHandler{service: service}
}

type createUserRequest struct {
	Name string `json:"name"`
}

func (h *UsersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	in := &createUserRequest{}
	if err := render.Decode(r, in); err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	user, err := h.service.CreateUser(in.Name)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
	}
	render.Respond(w, r, user)
}

type getUserResponse struct {
	User domain.User `json:"user"`
}

func (h *UsersHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userID := ctxKey.GetFormCtx(r, userIDCtxKey)
	user, err := h.service.GetUser(userID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, getUserResponse{User: *user})
}

func (h *UsersHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := ctxKey.GetFormCtx(r, userIDCtxKey)
	err := h.service.DeleteUser(userID)
	if err != nil {
		render.Status(r, resovleStatus(err))
		render.Respond(w, r, resp.ErrorResponse{Error: err.Error()})
		return
	}
	render.Respond(w, r, resp.OK())
}
