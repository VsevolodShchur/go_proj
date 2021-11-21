package domain

type UsersRepo interface {
	CreateUser(user *User) error
	GetUser(ID string) (*User, error)
	DeleteUser(ID string) error
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Notes []*Note `json:"notes"`
}
