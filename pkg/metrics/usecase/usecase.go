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
	var err error
	out := new(models.MetricsAggregated)

	out.PublicationsCount, err = u.publicationUC.GetPublicationsCount(ctx)
	if err != nil {
		return nil, err
	}

	ms, err := u.repo.GetMetrics(ctx, time.Now().AddDate(0, 0, -1))
	if err != nil {
		return nil, err
	}
	out.Last24HoursStats = ms

	return out, nil
}

//func (u MetricsUseCase) aggregateMetrics(ctx context.Context, ms []*models.Metrics) (*models.MetricsAggregated, error) {
//	aggregated := new(models.MetricsAggregated)
//	aggregated.Last24HoursStats = new(models.Metrics)
//
//	var err error
//	aggregated.PublicationsCount, err = u.publicationUC.GetPublicationsCount(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	// last 24 hours
//	fromTime := time.Now().AddDate(0, 0, -1)
//	aggregated.Last24HoursMax.UniqueVisitorsCount = getMaxUniqueUsers(ms, fromTime)
//
//	// last hour
//	fromTime = time.Now().Add(-time.Hour)
//	aggregated.Last24HoursMax.UniqueVisitorsCount = getMaxUniqueUsers(ms, fromTime)
//
//	return aggregated, nil
//}
//
//func getMaxUniqueUsers(ms []*models.Metrics, fromDate time.Time) int64 {
//	maxUnique := ms[0].UniqueVisitorsCount
//
//	for _, m := range ms {
//		if m.Timestamp.Unix() < fromDate.Unix() {
//			break
//		}
//
//		if m.UniqueVisitorsCount > maxUnique {
//			maxUnique = m.UniqueVisitorsCount
//		}
//	}
//
//	return maxUnique
//}
