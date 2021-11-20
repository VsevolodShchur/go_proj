package notes

import (
	"errors"
	"proj/internal/domain"
	mock_notes "proj/internal/service/notes/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type notesServiceMocks struct {
	repo *mock_notes.MockNotesRepo
}

var validUUID = "asd"

func newNotesServiceMocks(t *testing.T, setBehavior func(*notesServiceMocks)) *notesServiceMocks {
	c := gomock.NewController(t)
	mocks := &notesServiceMocks{
		repo: mock_notes.NewMockNotesRepo(c),
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
			inputUserID: validUUID,
			inputNoteID: validUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().GetNote(validUUID).Return(&domain.Note{
					ID:     validUUID,
					UserID: validUUID,
					Text:   "asd",
				}, nil)
				nsm.repo.EXPECT().UpdateNote(validUUID, "text").Return(nil)
			},
			expectedError: nil,
		},
		{
			name:        "when GetNote fails",
			inputUserID: validUUID,
			inputNoteID: validUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().GetNote(validUUID).Return(nil, errors.New("GetNote error"))
			},
			expectedError: errors.New("GetNote error"),
		},
		{
			name:        "when UpdateNote fails",
			inputUserID: validUUID,
			inputNoteID: validUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().GetNote(validUUID).Return(&domain.Note{
					ID:     validUUID,
					UserID: validUUID,
					Text:   "asd",
				}, nil)
				nsm.repo.EXPECT().UpdateNote(validUUID, "text").Return(errors.New("UpdateNote error"))
			},
			expectedError: errors.New("UpdateNote error"),
		},
		{
			name:        "user does not have this note",
			inputUserID: validUUID,
			inputNoteID: validUUID,
			inputText:   "text",
			mocksBehavior: func(nsm *notesServiceMocks) {
				nsm.repo.EXPECT().GetNote(validUUID).Return(&domain.Note{
					ID:     validUUID,
					UserID: "other uuid",
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

			err := s.UpdateNote(tc.inputUserID, tc.inputNoteID, tc.inputText)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
