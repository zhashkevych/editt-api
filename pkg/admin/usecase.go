package admin

import (
	"context"
	"edittapi/pkg/models"
)

type UseCase interface {
	GetMetrics(ctx context.Context) (*models.MetricsAggregated, error)
	GetFeedbacks(ctx context.Context) ([]*models.Feedback, error)

	GetAllPublications(ctx context.Context) ([]*models.Publication, error)
	RemovePublication(ctx context.Context, id string) error
}
