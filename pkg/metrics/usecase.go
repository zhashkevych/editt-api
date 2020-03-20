package metrics

import (
	"context"
	"edittapi/pkg/models"
)

type UseCase interface {
	SetMetrics(ctx context.Context, metrics *models.Metrics) error
	GetMetrics(ctx context.Context) ([]*models.Metrics, error)
}