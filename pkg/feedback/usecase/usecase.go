package usecase

import (
	"context"
	"edittapi/pkg/feedback"
	"edittapi/pkg/models"
)

type FeedbackUseCase struct {
	repo feedback.Repository
}

func NewFeedbackUseCase(repo feedback.Repository) *FeedbackUseCase {
	return &FeedbackUseCase{
		repo: repo,
	}
}

func (u *FeedbackUseCase) CreateFeedback(ctx context.Context, inp models.Feedback) error {
	return u.repo.Insert(ctx, inp)
}

func (u *FeedbackUseCase) GetAll(ctx context.Context) ([]*models.Feedback, error) {
	return u.repo.Get(ctx)
}



