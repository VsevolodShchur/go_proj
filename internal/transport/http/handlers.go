package http

import "proj/internal/service"

type Handlers struct {
	usersHandler *UsersHandler
	notesHandler *NotesHandler
}

func NewHandlers(s *service.Services) *Handlers {
	return &Handlers{
		usersHandler: NewUsersHandler(s.UserService),
		notesHandler: NewNotesHandler(s.NotesService),
	}
}
