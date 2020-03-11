package usecase

import (
	"context"
	"edittapi/application/profile/usecase"
	"edittapi/auth"
	"github.com/stretchr/testify/assert"
	"edittapi/auth/repository/mock"
	"edittapi/models"
	"testing"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)
	profileUC := new(usecase.ProfileUseCaseMock)

	uc := NewAuthUseCase(repo, profileUC, "salt", []byte("secret"), 86400)

	var (
		email    = "email"
		username = "user"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Email:    email,
			Username: username,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, auth.SignUpInput{email, auth.SignInInput{username, password}})
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, auth.SignInInput{username, password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
