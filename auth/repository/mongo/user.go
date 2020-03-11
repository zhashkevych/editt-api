package mongo

import (
	"context"
	"edittapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toUser(user)
	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User)

	// try to find with username
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err == nil {
		return toModel(user), nil
	}

	// try to find with email
	err = r.db.FindOne(ctx, bson.M{
		"email":    username,
		"password": password,
	}).Decode(user)

	if err == nil {
		return toModel(user), nil
	}

	return nil, err
}

func toUser(u *models.User) *User {
	return &User{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}
