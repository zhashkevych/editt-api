package metrics

import (
	"context"
	"edittapi/pkg/models"
	"time"
)

type Repository interface {
	SetMetrics(ctx context.Context, metrics models.Metrics) error
	GetMetrics(ctx context.Context, timeFrom time.Time) ([]*models.Metrics, error)
}
