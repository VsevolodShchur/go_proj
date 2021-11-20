package internal

//go:generate mockgen -source=service/notes/notes.go -destination=service/notes/mocks/notes.go
//go:generate mockgen -source=service/users/users.go -destination=service/users/mocks/users.go

//go:generate mockgen -source=transport/http/users.go -destination=transport/http/mocks/users.go
//go:generate mockgen -source=transport/http/notes.go -destination=transport/http/mocks/notes.go
