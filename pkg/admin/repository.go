package admin

import (
	"context"
	"edittapi/pkg/models"
)

type Repository interface {
	GetPublications(ctx context.Context) ([]*models.Publication, error)
	RemovePublication(ctx context.Context, id string) error
}
