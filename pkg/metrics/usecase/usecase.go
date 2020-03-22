package usecase

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
	"edittapi/pkg/publication"
)

type MetricsUseCase struct {
	repo          metrics.Repository
	publicationUC publication.UseCase
}

func NewMetricsUseCase(repo metrics.Repository, puc publication.UseCase) *MetricsUseCase {
	return &MetricsUseCase{
		repo:          repo,
		publicationUC: puc,
	}
}

func (u MetricsUseCase) SetMetrics(ctx context.Context, metrics models.Metrics) error {
	var err error

	metrics.PublicationsCount, err = u.publicationUC.GetPublicationsCount(ctx)
	if err != nil {
		return err
	}

	return u.repo.SetMetrics(ctx, metrics)
}

func (u MetricsUseCase) GetMetrics(ctx context.Context) ([]*models.Metrics, error) {
	return u.repo.GetMetrics(ctx)
}
