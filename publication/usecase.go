package publication

import (
	"context"
	"edittapi/models"
)

type PublicationUseCase interface {
	Publish(ctx context.Context, publication *models.Publication) error
	GetPublication(ctx context.Context, id string) (*models.Publication, error)
	GetUserPublications(ctx context.Context, user *models.User) ([]*models.Publication, error)
	UpdatePublication(ctx context.Context, publication *models.Publication) error
	DeletePublication(ctx context.Context, publication *models.Publication) error
}

type CommentUseCase interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	RecievePublicationComments(ctx context.Context, publication *models.Publication) ([]*models.Comment, error)
	DeleteComment(ctx context.Context, comment *models.Comment) error
}

type LikeUseCase interface {
	LikePublication(ctx context.Context, publication *models.Publication) error
}