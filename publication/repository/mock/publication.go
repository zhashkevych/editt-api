package mock

import (
	"context"
	"edittapi/models"
	"github.com/stretchr/testify/mock"
)

type PublicationRepoMock struct {
	mock.Mock
}

func (r *PublicationRepoMock) Create(ctx context.Context, publication models.Publication) error {
	args := r.Called(publication)

	return args.Error(0)
}

func (r *PublicationRepoMock) GetPopular(ctx context.Context, limit int64) ([]*models.Publication, error) {
	args := r.Called(limit)

	return args.Get(0).([]*models.Publication), args.Error(1)
}

func (r *PublicationRepoMock) GetLatest(ctx context.Context, limit int64) ([]*models.Publication, error) {
	args := r.Called(limit)

	return args.Get(0).([]*models.Publication), args.Error(1)
}

func (r *PublicationRepoMock) GetById(ctx context.Context, id string) (*models.Publication, error) {
	args := r.Called(id)

	return args.Get(0).(*models.Publication), args.Error(1)
}

func (r *PublicationRepoMock) IncrementReactions(ctx context.Context, id string) error {
	args := r.Called(id)

	return args.Error(0)
}

func (r *PublicationRepoMock) IncrementViews(ctx context.Context, id string) error {
	args := r.Called(id)

	return args.Error(0)
}