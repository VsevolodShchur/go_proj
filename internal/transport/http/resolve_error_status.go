package http

import (
	"errors"
	"net/http"
	notesRepo "proj/internal/repository/notes"
	"proj/internal/repository/users"
	"proj/internal/service/notes"
)

func resovleStatus(err error) int {
	if errors.Is(err, users.ErrUserNotFound) ||
		errors.Is(err, notesRepo.ErrNoteNotFound) ||
		errors.Is(err, notes.ErrUserNote) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
