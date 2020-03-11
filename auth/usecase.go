package auth

import (
	"context"
	"edittapi/models"
)

const CtxUserKey = "user"

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	SignInInput
}

type UseCase interface {
	SignUp(ctx context.Context, inp SignUpInput) error
	SignIn(ctx context.Context, inp SignInInput) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
	AttachProfile(ctx context.Context, user *models.User) error
}
