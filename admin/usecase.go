package admin

import (
	"context"
	"edittapi/models"
)

type UseCase interface {
	GetMetrics(ctx context.Context) (*models.Metrics, error)

	GetAllPublications(ctx context.Context) ([]*models.Publication, error)
	RemovePublication(ctx context.Context, id string) error
}
