package usecase

import (
	"context"
	"edittapi/application/profile"
	"github.com/stretchr/testify/mock"
	"edittapi/models"
)

type ProfileUseCaseMock struct {
	mock.Mock
}

func (p ProfileUseCaseMock) SetUserProfile(ctx context.Context, user *models.User) error {
	args := p.Called(user)

	return args.Error(1)
}

func (p ProfileUseCaseMock) CreateProfile(ctx context.Context, user *models.User, inp profile.ProfileInput) (*models.Profile, error) {
	args := p.Called(user, user, inp)

	return args.Get(0).(*models.Profile), args.Error(1)
}

func (p ProfileUseCaseMock) UpdateProfile(ctx context.Context, user *models.User, inp profile.ProfileInput) error {
	args := p.Called(user, user, inp)

	return args.Error(0)
}

func (p ProfileUseCaseMock) DeleteProfile(ctx context.Context, user *models.User) error {
	args := p.Called(user, user)

	return args.Error(0)
}
