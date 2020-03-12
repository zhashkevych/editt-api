package usecase

import (
	"context"
	"edittapi/models"
	"github.com/stretchr/testify/mock"
)

type BookmarkUseCaseMock struct {
	mock.Mock
}

func (m BookmarkUseCaseMock) CreateBookmark(ctx context.Context, user *models.User, url, title string) error {
	args := m.Called(user, url, title)

	return args.Error(0)
}

func (m BookmarkUseCaseMock) GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	args := m.Called(user)

	return args.Get(0).([]*models.Bookmark), args.Error(1)
}

func (m BookmarkUseCaseMock) DeleteBookmark(ctx context.Context, user *models.User, id string) error {
	args := m.Called(user, id)

	return args.Error(0)
}
