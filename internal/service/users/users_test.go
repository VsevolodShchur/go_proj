package users

import (
	"errors"
	"proj/internal/domain"
	mock_notes "proj/internal/service/notes/mocks"
	mock_users "proj/internal/service/users/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var validUUID = "asd"

type usersServiceMocks struct {
	repo      *mock_users.MockUsersRepo
	notesRepo *mock_notes.MockNotesRepo
}

func newUsersServiceMocks(t *testing.T, setBehavior func(*usersServiceMocks)) *usersServiceMocks {
	c := gomock.NewController(t)
	mocks := &usersServiceMocks{
		repo:      mock_users.NewMockUsersRepo(c),
		notesRepo: mock_notes.NewMockNotesRepo(c),
	}
	setBehavior(mocks)
	return mocks
}

func TestUsersService_GetUser(t *testing.T) {
	tcs := []struct {
		name          string
		input         string
		mocksBehavior func(*usersServiceMocks)
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:  "ok",
			input: validUUID,
			mocksBehavior: func(usm *usersServiceMocks) {
				usm.repo.EXPECT().GetUser(validUUID).Return(&domain.User{
					ID:   validUUID,
					Name: "",
				}, nil)
				usm.notesRepo.EXPECT().ListUserNotes(validUUID).Return([]*domain.Note{
					{ID: validUUID, Text: "asd"},
				}, nil)
			},
			expectedUser: &domain.User{ID: validUUID, Name: "", Notes: []*domain.Note{
				{ID: validUUID, Text: "asd"},
			}},
			expectedError: nil,
		},
		{
			name:  "when GetUser fails",
			input: validUUID,
			mocksBehavior: func(usm *usersServiceMocks) {
				usm.repo.EXPECT().GetUser(validUUID).Return(nil, errors.New("GetUser error"))
			},
			expectedUser:  nil,
			expectedError: errors.New("GetUser error"),
		},
		{
			name:  "when ListUserNotes fails",
			input: validUUID,
			mocksBehavior: func(usm *usersServiceMocks) {
				usm.repo.EXPECT().GetUser(validUUID).Return(&domain.User{
					ID:   validUUID,
					Name: "",
				}, nil)
				usm.notesRepo.EXPECT().ListUserNotes(validUUID).Return(nil, errors.New("ListUserNotes error"))
			},
			expectedUser:  nil,
			expectedError: errors.New("ListUserNotes error"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mocks := newUsersServiceMocks(t, tc.mocksBehavior)
			s := NewUserService(mocks.repo, mocks.notesRepo)

			user, err := s.GetUser(tc.input)

			assert.Equal(t, tc.expectedUser, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
