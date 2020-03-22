package usecase

import (
	"context"
	"edittapi/pkg/models"
	"edittapi/pkg/publication"
	"github.com/microcosm-cc/bluemonday"
	"strings"
	"time"
)

const (
	averageReadingSpeed = 200 // Words per minute
	bodyLengthLimit     = 50000
)

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

	if err := validateBodyLength(publication.Body); err != nil {
		return err
	}

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

func (p PublicationUseCase) IncrementReactions(ctx context.Context, id string) error {
	return p.repo.IncrementReactions(ctx, id)
}

func (p PublicationUseCase) IncrementViews(ctx context.Context, id string) error {
	return p.repo.IncrementViews(ctx, id)
}

func (p PublicationUseCase) GetPublications(ctx context.Context) ([]*models.Publication, error) {
	return p.repo.GetPublications(ctx)
}

func (p PublicationUseCase) GetPublicationsCount(ctx context.Context) (int64, error) {
	return p.repo.GetPublicationsCount(ctx)
}

func (p PublicationUseCase) RemovePublication(ctx context.Context, id string) error {
	return p.repo.RemovePublication(ctx, id)
}

func estimateReadingTime(text string) int32 {
	wordsCount := len(strings.Split(text, " "))

	readingTime := int32(wordsCount / averageReadingSpeed)
	if readingTime == 0 {
		readingTime = 1
	}

	return readingTime
}

func validateBodyLength(text string) error {
	if len(strings.Split(text, "")) > bodyLengthLimit {
		return publication.ErrWordsLimitExceeded
	}
	return nil
}
