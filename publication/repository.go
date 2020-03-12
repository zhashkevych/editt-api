package publication

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	Create(ctx context.Context, publication models.Publication) error
	GetPopular(ctx context.Context) ([]*models.Publication, error)
	GetLatest(ctx context.Context) ([]*models.Publication, error)
}
