package usecase

import (
	"context"
	"edittapi/application/profile"
	"edittapi/models"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProfileUseCase struct {
	repo profile.Repository
}

func NewProfileUseCase(repo profile.Repository) *ProfileUseCase {
	return &ProfileUseCase{
		repo: repo,
	}
}

func (p ProfileUseCase) SetUserProfile(ctx context.Context, user *models.User) error {
	profile, err := p.repo.Get(ctx, user)
	if err == nil {
		user.Profile = profile
		return nil
	}

	profile = &models.Profile{
		ProfileIsSet: false,
	}

	if _, err := p.repo.Create(ctx, profile); err != nil {
		log.Errorf("Failed to create profile for user %s: %s", user.ID, err.Error())
		return err
	}

	user.Profile = profile
	return nil
}

func (p ProfileUseCase) UpdateProfile(ctx context.Context, user *models.User, inp profile.ProfileInput) error {
	profile := toProfile(user, inp)

	return p.repo.Update(ctx, profile)
}

func (p ProfileUseCase) DeleteProfile(ctx context.Context, user *models.User) error {
	return p.repo.Delete(ctx, &models.Profile{})
}

func toProfile(user *models.User, inp profile.ProfileInput) *models.Profile {
	return &models.Profile{
		UserID:       user.ID,
		ProfileIsSet: true,
		FirstName:    inp.FirstName,
		LastName:     inp.LastName,
		Bio:          inp.Bio,
		ProfileImage: inp.ImageLink,
		CreatedAt:    time.Now(),
		Interests:    inp.Interests,
	}
}
