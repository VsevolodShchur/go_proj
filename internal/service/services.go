package service

import (
	"proj/internal/repository"
	"proj/internal/service/notes"
	"proj/internal/service/users"
)

type Services struct {
	UserService  *users.UserService
	NotesService *notes.NotesService
}

func Init(repos *repository.Repositories) *Services {
	return &Services{
		UserService:  users.NewUserService(repos.UsersRepo, repos.NotesRepo),
		NotesService: notes.NewNotesService(repos.NotesRepo),
	}
}
