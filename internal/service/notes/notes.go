package notes

import (
	"errors"
	"proj/internal/domain"

	"github.com/google/uuid"
)

var ErrUserNote = errors.New("user does not have this note")

type NotesRepo interface {
	CreateNote(note *domain.Note) error
	GetNote(ID string) (*domain.Note, error)
	UpdateNote(ID string, text string) error
	DeleteNote(ID string) error
	ListUserNotes(userID string) ([]*domain.Note, error)
}

type NotesService struct {
	repo NotesRepo
}

func NewNotesService(repo NotesRepo) *NotesService {
	return &NotesService{
		repo: repo,
	}
}

func (s *NotesService) CreateNote(text string, userID string) (*domain.Note, error) {
	note := &domain.Note{
		ID:     uuid.NewString(),
		Text:   text,
		UserID: userID,
	}
	if err := s.repo.CreateNote(note); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *NotesService) GetNote(ID string) (*domain.Note, error) {
	return s.repo.GetNote(ID)
}

func (s *NotesService) checkUserNote(userID string, noteID string) error {
	note, err := s.repo.GetNote(noteID)
	if err != nil {
		return err
	}
	if note.UserID != userID {
		return ErrUserNote
	}
	return nil
}

func (s *NotesService) UpdateNote(userID string, noteID string, text string) error {
	if err := s.checkUserNote(userID, noteID); err != nil {
		return err
	}
	return s.repo.UpdateNote(noteID, text)
}

func (s *NotesService) DeleteNote(userID string, noteID string) error {
	if err := s.checkUserNote(userID, noteID); err != nil {
		return err
	}
	return s.repo.DeleteNote(noteID)
}

func (s *NotesService) ListUserNotes(userID string) ([]*domain.Note, error) {
	return s.repo.ListUserNotes(userID)
}
