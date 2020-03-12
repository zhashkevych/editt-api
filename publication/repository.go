package publication

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	Create(ctx context.Context, publication models.Publication) error
	GetPopular(ctx context.Context, limit int64) ([]*models.Publication, error)
	GetLatest(ctx context.Context, limit int64) ([]*models.Publication, error)
}
