package usecase

import (
	"context"
	"edittapi/pkg/models"

	"github.com/stretchr/testify/mock"
)

type PublicationUseCaseMock struct {
	mock.Mock
}

func (p *PublicationUseCaseMock) Publish(ctx context.Context, publication models.Publication) error {
	args := p.Called(publication)

	return args.Error(0)
}

func (p *PublicationUseCaseMock) GetPopularPublications(ctx context.Context, limit int64) ([]*models.Publication, error) {
	args := p.Called(limit)

	return args.Get(0).([]*models.Publication), args.Error(1)
}

func (p *PublicationUseCaseMock) GetLatestPublications(ctx context.Context, limit int64) ([]*models.Publication, error) {
	args := p.Called(limit)

	return args.Get(0).([]*models.Publication), args.Error(1)
}

func (p *PublicationUseCaseMock) GetById(ctx context.Context, id string) (*models.Publication, error) {
	args := p.Called(id)

	return args.Get(0).(*models.Publication), args.Error(1)
}

func (p *PublicationUseCaseMock) IncrementReactions(ctx context.Context, id string) error {
	args := p.Called(id)

	return args.Error(0)
}

func (p *PublicationUseCaseMock) IncrementViews(ctx context.Context, id string) error {
	args := p.Called(id)

	return args.Error(0)
}
