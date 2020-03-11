package profile

import (
	"context"
	"edittapi/models"
)

type ProfileInput struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Nickname  string   `json:"nickname"`
	Bio       string   `json:"bio"`
	ImageLink string   `json:"imageLink"`
	Interests []string `json:"interests"`
}

type UseCase interface {
	UpdateProfile(ctx context.Context, user *models.User, inp ProfileInput) error
	DeleteProfile(ctx context.Context, user *models.User) error
	SetUserProfile(ctx context.Context, user *models.User) error
}
