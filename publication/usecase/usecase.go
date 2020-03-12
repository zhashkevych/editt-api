package usecase

import (
	"context"
	"edittapi/models"
	"edittapi/publication"
	"github.com/microcosm-cc/bluemonday"
	"strings"
	"time"
)

const averageReadingSpeed = 200 // Words per minute

type PublicationUseCase struct {
	repo   publication.Repository
	policy *bluemonday.Policy
}

func NewPublicationUseCase(repo publication.Repository) *PublicationUseCase {
	return &PublicationUseCase{
		repo:   repo,
		policy: bluemonday.UGCPolicy(),
	}
}

func (p PublicationUseCase) Publish(ctx context.Context, publication models.Publication) error {
	publication.PublishedAt = time.Now()
	publication.Body = p.policy.Sanitize(publication.Body)
	publication.ReadingTime = estimateReadingTime(publication.Body)

	return p.repo.Create(ctx, publication)
}

func (p PublicationUseCase) GetPopularPublications(ctx context.Context, limit int64) ([]*models.Publication, error) {
	return p.repo.GetPopular(ctx, limit)
}

func (p PublicationUseCase) GetLatestPublications(ctx context.Context, limit int64) ([]*models.Publication, error) {
	return p.repo.GetLatest(ctx, limit)
}

func (p PublicationUseCase) GetById(ctx context.Context, id string) (*models.Publication, error) {
	return p.repo.GetById(ctx, id)
}

func (p PublicationUseCase) IncrementClaps(ctx context.Context, id string) error {
	return p.repo.IncrementClaps(ctx, id)
}

func (p PublicationUseCase) IncrementViews(ctx context.Context, id string) error {
	return p.repo.IncrementViews(ctx, id)
}

func estimateReadingTime(text string) int32 {
	wordsCount := len(strings.Split(text, ""))

	return int32(wordsCount / averageReadingSpeed)
}
