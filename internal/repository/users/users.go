package users

import (
	"errors"
	"proj/internal/domain"
)

var ErrUserNotFound = errors.New("user not found")

type UsersRepo struct {
	users map[string]*domain.User
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{users: make(map[string]*domain.User)}
}

func (r *UsersRepo) CreateUser(user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UsersRepo) GetUser(ID string) (*domain.User, error) {
	user, ok := r.users[ID]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *UsersRepo) DeleteUser(ID string) error {
	if _, ok := r.users[ID]; !ok {
		return ErrUserNotFound
	}
	delete(r.users, ID)
	return nil
}
