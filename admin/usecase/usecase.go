package usecase

import (
	"context"
	"edittapi/admin"
	"edittapi/metrics"
	"edittapi/models"
)

type AdminUseCase struct {
	metricsRepo     metrics.Repository
	publicationRepo admin.Repository
}

func NewAdminUseCase(mr metrics.Repository, pr admin.Repository) *AdminUseCase {
	return &AdminUseCase{
		metricsRepo:     mr,
		publicationRepo: pr,
	}
}

func (u AdminUseCase) GetMetrics(ctx context.Context) (*models.Metrics, error) {
	return u.metricsRepo.GetMetrics(ctx)
}

func (u AdminUseCase) GetAllPublications(ctx context.Context) ([]*models.Publication, error) {
	return u.publicationRepo.GetPublications(ctx)
}

func (u AdminUseCase) RemovePublication(ctx context.Context, id string) error {
	return u.publicationRepo.RemovePublication(ctx, id)
}
