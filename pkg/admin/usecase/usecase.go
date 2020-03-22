package usecase

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
	"edittapi/pkg/publication"
)

type AdminUseCase struct {
	metricsUseCase     metrics.UseCase
	publicationUseCase publication.UseCase
}

func NewAdminUseCase(mr metrics.UseCase, pr publication.UseCase) *AdminUseCase {
	return &AdminUseCase{
		metricsUseCase:     mr,
		publicationUseCase: pr,
	}
}

func (u AdminUseCase) GetMetrics(ctx context.Context) ([]*models.Metrics, error) {
	return u.metricsUseCase.GetMetrics(ctx)
}

func (u AdminUseCase) GetAllPublications(ctx context.Context) ([]*models.Publication, error) {
	return u.publicationUseCase.GetPublications(ctx)
}

func (u AdminUseCase) RemovePublication(ctx context.Context, id string) error {
	return u.publicationUseCase.RemovePublication(ctx, id)
}
