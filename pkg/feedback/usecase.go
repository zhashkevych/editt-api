package feedback

import (
	"context"
	"edittapi/pkg/models"
)

type UseCase interface {
	CreateFeedback(ctx context.Context, inp models.Feedback) error
	GetAll(ctx context.Context) ([]*models.Feedback, error)
}
