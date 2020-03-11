package profile

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	Get(ctx context.Context, user *models.User) (*models.Profile, error)
	Create(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	Update(ctx context.Context, profile *models.Profile) error
	Delete(ctx context.Context, profile *models.Profile) error
}
