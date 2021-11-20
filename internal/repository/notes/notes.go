package notes

import (
	"errors"
	"proj/internal/domain"
)

var ErrNoteNotFound = errors.New("note not found")

type NotesRepo struct {
	notes map[string]*domain.Note
}

func NewNotesRepo() *NotesRepo {
	return &NotesRepo{notes: make(map[string]*domain.Note)}
}

func (r *NotesRepo) CreateNote(note *domain.Note) error {
	r.notes[note.ID] = note
	return nil
}

func (r *NotesRepo) GetNote(ID string) (*domain.Note, error) {
	note, ok := r.notes[ID]
	if !ok {
		return nil, ErrNoteNotFound
	}
	return note, nil
}

func (r *NotesRepo) UpdateNote(ID string, text string) error {
	note, ok := r.notes[ID]
	if !ok {
		return ErrNoteNotFound
	}
	note.Text = text
	r.notes[ID] = note
	return nil
}

func (r *NotesRepo) DeleteNote(ID string) error {
	if _, ok := r.notes[ID]; !ok {
		return ErrNoteNotFound
	}
	delete(r.notes, ID)
	return nil
}

func (r *NotesRepo) ListUserNotes(userID string) ([]*domain.Note, error) {
	notes := make([]*domain.Note, 0)
	for _, note := range r.notes {
		if note.UserID == userID {
			notes = append(notes, note)
		}
	}
	return notes, nil
}
