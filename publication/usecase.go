package publication

import (
	"context"
	"edittapi/models"
)

type UseCase interface {
	Publish(ctx context.Context, publication *models.Publication) error
	GetPopularPublications(ctx context.Context) ([]*models.Publication, error)
	GetLatestPublications(ctx context.Context) ([]*models.Publication, error)
}