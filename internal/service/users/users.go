package users

import (
	"proj/internal/domain"

	"github.com/google/uuid"
)

type notesRepo interface {
	ListUserNotes(userID string) ([]*domain.Note, error)
	DeleteUserNotes(userID string) error
}

type UserService struct {
	repo      domain.UsersRepo
	notesRepo notesRepo
}

func NewUserService(repo domain.UsersRepo, notesRepo notesRepo) *UserService {
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
	if err := s.notesRepo.DeleteUserNotes(ID); err != nil {
		return err
	}
	return s.repo.DeleteUser(ID)
}
