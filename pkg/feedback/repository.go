package feedback

import (
	"context"
	"edittapi/pkg/models"
)

type Repository interface {
	Insert(ctx context.Context, inp models.Feedback) error
	Get(ctx context.Context) ([]*models.Feedback, error)
}
