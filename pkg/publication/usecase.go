package publication

import (
	"context"
	"edittapi/pkg/models"
)

type UseCase interface {
	Publish(ctx context.Context, publication models.Publication) error

	GetPopularPublications(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetLatestPublications(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetById(ctx context.Context, id string) (*models.Publication, error)

	IncrementReactions(ctx context.Context, id string) error
	IncrementViews(ctx context.Context, id string) error

	GetPublications(ctx context.Context) ([]*models.Publication, error)
	GetPublicationsCount(ctx context.Context) (int64, error)

	RemovePublication(ctx context.Context, id string) error
}
