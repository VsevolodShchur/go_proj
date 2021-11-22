package notes

import (
	"errors"
	"proj/internal/domain"
	mock_domain "proj/internal/domain/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type notesServiceMocks struct {
	repo *mock_domain.MockNotesRepo
}

var (
	userUUID = "43b6408d-e195-42e5-87ff-fe97c24ea080"
	noteUUID = "c37c0de8-a784-4ec1-8ca9-598535661c2c"
)

func newNotesServiceMocks(t *testing.T,
	setBehavior func(*notesServiceMocks)) *notesServiceMocks {
	c := gomock.NewController(t)
	mocks := &notesServiceMocks{
		repo: mock_domain.NewMockNotesRepo(c),
	}
	setBehavior(mocks)
	return mocks
}

func TestNotesService_UpdateNote(t *testing.T) {
	tcs := []struct {
		name          string
		inputUserID   string
		inputNoteID   string
		inputText     string
		mocksBehavior func(*notesServiceMocks)
		expectedError error
	}{
		{
			name:        "ok",
			inputUserID: userUUID,
			inputNoteID: noteUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().
					GetNote(noteUUID).
					Return(
						&domain.Note{
							ID:     noteUUID,
							UserID: userUUID,
							Text:   "asd",
						}, nil)
				nsm.repo.EXPECT().
					UpdateNote(noteUUID, "text").
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:        "when GetNote fails",
			inputUserID: userUUID,
			inputNoteID: noteUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().
					GetNote(noteUUID).
					Return(nil, errors.New("GetNote error"))
			},
			expectedError: errors.New("GetNote error"),
		},
		{
			name:        "when UpdateNote fails",
			inputUserID: userUUID,
			inputNoteID: noteUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().
					GetNote(noteUUID).
					Return(
						&domain.Note{
							ID:     noteUUID,
							UserID: userUUID,
							Text:   "asd",
						}, nil)
				nsm.repo.EXPECT().
					UpdateNote(noteUUID, "text").
					Return(errors.New("UpdateNote error"))
			},
			expectedError: errors.New("UpdateNote error"),
		},
		{
			name:        "user does not have this note",
			inputUserID: userUUID,
			inputNoteID: noteUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().
					GetNote(noteUUID).
					Return(
						&domain.Note{
							ID:     noteUUID,
							UserID: "437e7406-76c7-4eff-af14-e5f42ed7c793",
							Text:   "asd",
						}, nil)
			},
			expectedError: ErrUserNote,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mocks := newNotesServiceMocks(t, tc.mocksBehavior)
			s := NewNotesService(mocks.repo)

			err := s.UpdateNote(tc.inputUserID, tc.inputNoteID,
				tc.inputText)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
