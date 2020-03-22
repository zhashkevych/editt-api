package usecase

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
	"edittapi/pkg/publication"
	"time"
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
	return u.repo.SetMetrics(ctx, metrics)
}

func (u MetricsUseCase) GetMetrics(ctx context.Context) (*models.MetricsAggregated, error) {
	ms, err := u.repo.GetMetrics(ctx)
	if err != nil {
		return nil, err
	}

	return u.aggregateMetrics(ctx, ms)
}

func (u MetricsUseCase) aggregateMetrics(ctx context.Context, ms []*models.Metrics) (*models.MetricsAggregated, error) {
	aggregated := new(models.MetricsAggregated)
	aggregated.Last24Hours = new(models.Metrics)
	aggregated.LastHour = new(models.Metrics)

	var err error
	aggregated.PublicationsCount, err = u.publicationUC.GetPublicationsCount(ctx)
	if err != nil {
		return nil, err
	}

	// last 24 hours
	fromTime := time.Now().AddDate(0, 0, -1)
	aggregated.Last24Hours.UniqueVisitorsCount = getAverageUniqueUsers(ms, fromTime)

	// last hour
	fromTime = time.Now().Add(-time.Hour)
	aggregated.Last24Hours.UniqueVisitorsCount = getAverageUniqueUsers(ms, fromTime)

	return aggregated, nil
}

func getAverageUniqueUsers(ms []*models.Metrics, fromDate time.Time) int64 {
	uniqueAvg := int64(0)

	for _, m := range ms {
		if m.Timestamp.Unix() < fromDate.Unix() {
			break
		}

		uniqueAvg += m.UniqueVisitorsCount
	}

	uniqueAvg /= int64(len(ms))

	return uniqueAvg
}
