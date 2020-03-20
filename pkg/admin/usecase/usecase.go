package usecase

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
	"edittapi/pkg/publication"
)

type AdminUseCase struct {
	metricsRepo     metrics.Repository
	publicationRepo publication.Repository
}

func NewAdminUseCase(mr metrics.Repository, pr publication.Repository) *AdminUseCase {
	return &AdminUseCase{
		metricsRepo:     mr,
		publicationRepo: pr,
	}
}

func (u AdminUseCase) GetMetrics(ctx context.Context) ([]*models.Metrics, error) {
	return u.metricsRepo.GetMetrics(ctx)
}

func (u AdminUseCase) GetAllPublications(ctx context.Context) ([]*models.Publication, error) {
	return u.publicationRepo.GetPublications(ctx)
}

func (u AdminUseCase) RemovePublication(ctx context.Context, id string) error {
	return u.publicationRepo.RemovePublication(ctx, id)
}
