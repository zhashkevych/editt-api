package publication

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	Create(ctx context.Context, publication models.Publication) error

	GetPopular(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetLatest(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetById(ctx context.Context, id string) (*models.Publication, error)

	IncrementReactions(ctx context.Context, id string) error
	IncrementViews(ctx context.Context, id string) error
}
