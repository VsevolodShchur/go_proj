package internal

//go:generate mockgen -source=domain/note.go -destination=domain/mocks/note.go
//go:generate mockgen -source=domain/user.go -destination=domain/mocks/user.go

//go:generate mockgen -source=transport/http/users.go -destination=transport/http/mocks/users.go
//go:generate mockgen -source=transport/http/notes.go -destination=transport/http/mocks/notes.go
