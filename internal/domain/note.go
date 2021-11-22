package domain

type NotesRepo interface {
	CreateNote(note *Note) error
	GetNote(ID string) (*Note, error)
	UpdateNote(ID string, text string) error
	DeleteNote(ID string) error
	ListUserNotes(userID string) ([]*Note, error)
	DeleteUserNotes(userID string) error
}

type Note struct {
	ID     string `json:"id"`
	UserID string `json:"-"`
	Text   string `json:"text"`
}
