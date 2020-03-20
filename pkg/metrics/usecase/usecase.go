package usecase

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
)

type MetricsUseCase struct {
	repo metrics.Repository
}

func NewMetricsUseCase(repo metrics.Repository) *MetricsUseCase {
	return &MetricsUseCase{
		repo: repo,
	}
}

func (u MetricsUseCase) SetMetrics(ctx context.Context, metrics *models.Metrics) error {
	return u.repo.SetMetrics(ctx, metrics)
}

func (u MetricsUseCase) GetMetrics(ctx context.Context) ([]*models.Metrics, error) {
	return u.repo.GetMetrics(ctx)
}

