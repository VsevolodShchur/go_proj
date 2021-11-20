package users

import (
	"proj/internal/domain"

	"github.com/google/uuid"
)

type UsersRepo interface {
	CreateUser(user *domain.User) error
	GetUser(ID string) (*domain.User, error)
	DeleteUser(ID string) error
}

type NotesRepo interface {
	ListUserNotes(userID string) ([]*domain.Note, error)
}

type UserService struct {
	repo      UsersRepo
	notesRepo NotesRepo
}

func NewUserService(repo UsersRepo, notesRepo NotesRepo) *UserService {
	return &UserService{
		repo:      repo,
		notesRepo: notesRepo,
	}
}

func (s *UserService) CreateUser(name string) (*domain.User, error) {
	user := &domain.User{
		ID:   uuid.NewString(),
		Name: name,
	}
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(ID string) (*domain.User, error) {
	user, err := s.repo.GetUser(ID)
	if err != nil {
		return nil, err
	}
	notes, err := s.notesRepo.ListUserNotes(ID)
	if err != nil {
		return nil, err
	}
	user.Notes = notes
	return user, nil
}

func (s *UserService) DeleteUser(ID string) error {
	return s.repo.DeleteUser(ID)
}
