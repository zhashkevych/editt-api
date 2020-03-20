package metrics

import (
	"context"
	"edittapi/pkg/models"
)

type Repository interface {
	SetMetrics(ctx context.Context, metrics models.Metrics) error
	GetMetrics(ctx context.Context) ([]*models.Metrics, error)
}
