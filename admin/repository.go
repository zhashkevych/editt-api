package admin

import (
	"context"
	"edittapi/models"
)

type Repository interface {
	GetPublications(ctx context.Context) ([]*models.Publication, error)
	RemovePublication(ctx context.Context, id string) error
}
