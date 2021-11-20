package repository

import (
	"proj/internal/repository/notes"
	"proj/internal/repository/users"
)

type Repositories struct {
	UsersRepo *users.UsersRepo
	NotesRepo *notes.NotesRepo
}

func Init() *Repositories {
	return &Repositories{
		UsersRepo: users.NewUsersRepo(),
		NotesRepo: notes.NewNotesRepo(),
	}
}
