package usecase

import (
	"context"
	"edittapi/models"
	"github.com/stretchr/testify/mock"
)

type PublicationUseCaseMock struct {
	mock.Mock
}

func (p PublicationUseCaseMock) Publish(ctx context.Context, publication models.Publication) error {
	args := p.Called(publication)

	return args.Error(0)
}

func (p PublicationUseCaseMock) GetPopularPublications(ctx context.Context) ([]*models.Publication, error) {
	args := p.Called(ctx)

	return args.Get(0).([]*models.Publication), args.Error(1)
}

func (p PublicationUseCaseMock) GetLatestPublications(ctx context.Context) ([]*models.Publication, error) {
	args := p.Called(ctx)

	return args.Get(0).([]*models.Publication), args.Error(1)
}
