package publication

import (
	"context"
	"edittapi/models"
)

type UseCase interface {
	Publish(ctx context.Context, publication models.Publication) error

	GetPopularPublications(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetLatestPublications(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetById(ctx context.Context, id string) (*models.Publication, error)

	IncrementClaps(ctx context.Context, id string) error
	IncrementViews(ctx context.Context, id string) error
}
