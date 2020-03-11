package mongo

import (
	"context"
	"edittapi/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Profile struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty"`
	UserID            primitive.ObjectID   `bson:"userId"`
	ProfileIsSet      bool                 `bson:"profileIsSet"`
	FirstName         string               `bson:"firstName"`
	LastName          string               `bson:"lastName"`
	Bio               string               `bson:"bio"`
	Followers         []primitive.ObjectID `bson:"followers"`
	Following         []primitive.ObjectID `bson:"following"`
	SavedPublications []primitive.ObjectID `bson:"savedPublications"`
	LikedPublications []primitive.ObjectID `bson:"likedPublications"`
	ProfileImage      string               `bson:"profileImage"`
	CreatedAt         time.Time            `bson:"createdAt"`
	Membership        primitive.ObjectID   `bson:"membershipId"`
	Interests         []string             `bson:"interests"`
}

type Membership struct {
	ID         primitive.ObjectID    `bson:"_id,omitempty"`
	UserID     primitive.ObjectID    `bson:"userId"`
	Type       models.MembershipType `bson:"type"`
	isExpiried bool                  `bson:"isExpiried"`
	StartedAt  time.Time             `bson:"startedAt"`
	ExpiresAt  time.Time             `bson:"expiresAt"`
}

type ProfileRepository struct {
	db *mongo.Collection
}

func NewProfileRepository(db *mongo.Database, collection string) *ProfileRepository {
	return &ProfileRepository{
		db: db.Collection(collection),
	}
}

func (r ProfileRepository) Get(ctx context.Context, user *models.User) (*models.Profile, error) {
	profile := new(Profile)

	uid, _ := primitive.ObjectIDFromHex(user.ID)
	if err := r.db.FindOne(ctx, bson.M{"userId": uid}).Decode(profile); err != nil {
		return nil, err
	}

	return toModel(profile), nil
}

func (r ProfileRepository) Create(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	model := toProfile(profile)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	profile.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return profile, nil
}

func (r ProfileRepository) Update(ctx context.Context, profile *models.Profile) error {
	updateProfile := toProfile(profile)

	_, err := r.db.UpdateOne(ctx, bson.M{"userId": updateProfile.UserID}, updateProfile)
	if err != nil {
		log.Errorf("error occured while updating profile: %s", err.Error())
		return err
	}

	return nil
}

func (r ProfileRepository) Delete(ctx context.Context, profile *models.Profile) error {
	objID, _ := primitive.ObjectIDFromHex(profile.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func toProfile(profile *models.Profile) *Profile {
	uid, _ := primitive.ObjectIDFromHex(profile.UserID)

	return &Profile{
		UserID:       uid,
		ProfileIsSet: true,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		Bio:          profile.Bio,
		ProfileImage: profile.ProfileImage,
		CreatedAt:    time.Now(),
		Interests:    profile.Interests,
	}
}

func toModel(p *Profile) *models.Profile {
	return &models.Profile{
		ID:           p.ID.Hex(),
		UserID:       p.UserID.Hex(),
		ProfileIsSet: p.ProfileIsSet,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Bio:          p.Bio,
		Followers:    toStringArray(p.Followers),
		Following:    toStringArray(p.Following),
		Saved:        toStringArray(p.SavedPublications),
		Liked:        toStringArray(p.LikedPublications),
		ProfileImage: p.ProfileImage,
		CreatedAt:    p.CreatedAt,
		Interests:    p.Interests,
	}
}

func toStringArray(objIds []primitive.ObjectID) []string {
	ids := make([]string, len(objIds))

	for i := range objIds {
		ids[i] = objIds[i].Hex()
	}

	return ids
}
