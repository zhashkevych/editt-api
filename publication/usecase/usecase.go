package usecase

import (
	"context"
	"edittapi/models"
	"edittapi/publication"
)

type PublicationUseCase struct {
	repo publication.Repository
}

func NewPublicationUseCase(repo publication.Repository) *PublicationUseCase {
	return &PublicationUseCase{
		repo: repo,
	}
}

func (p PublicationUseCase) Publish(ctx context.Context, publication models.Publication) error {
	return p.repo.Create(ctx, publication)
}

func (p PublicationUseCase) GetPopularPublications(ctx context.Context) ([]*models.Publication, error) {
	return p.repo.GetPopular(ctx)
}

func (p PublicationUseCase) GetLatestPublications(ctx context.Context) ([]*models.Publication, error) {
	return p.repo.GetLatest(ctx)
}
