package metrics

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	GetMetrics(ctx context.Context) (*models.Metrics, error)
	SetMetrics(ctx context.Context) (*models.Metrics, error)
}
