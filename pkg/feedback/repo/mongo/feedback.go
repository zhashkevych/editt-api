package mongo

import (
	"context"
	"edittapi/pkg/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Feedback struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Score    int32              `bson:"score"`
	Features []int32            `bson:"features"`
}

type FeedbackRepository struct {
	db *mongo.Collection
}

func NewFeedbackRepository(db *mongo.Database, collection string) *FeedbackRepository {
	return &FeedbackRepository{
		db: db.Collection(collection),
	}
}

func (r *FeedbackRepository) Insert(ctx context.Context, inp models.Feedback) error {
	f := toFeedback(inp)

	_, err := r.db.InsertOne(ctx, f)
	if err != nil {
		log.Errorf("Publication Repo: error occured on inserting publication: %s", err.Error())
		return err
	}

	return nil
}

func (r *FeedbackRepository) Get(ctx context.Context) ([]*models.Feedback, error) {
	var feedbacks []*Feedback

	cur, err := r.db.Find(ctx, bson.D{{}})
	if err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	for cur.Next(ctx) {
		var elem Feedback
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("Publication Repo: error occured while decoding popular publications: %s", err.Error())
			return nil, err
		}

		feedbacks = append(feedbacks, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	if err := cur.Close(ctx); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	return toModels(feedbacks), nil
}

func toFeedback(f models.Feedback) *Feedback {
	return &Feedback{
		Score:    f.Score,
		Features: f.Features,
	}
}

func toModels(fs []*Feedback) []*models.Feedback {
	out := make([]*models.Feedback, len(fs))
	for i := range fs {
		out[i] = toModel(fs[i])
	}

	return out
}

func toModel(f *Feedback) *models.Feedback {
	return &models.Feedback{
		Score:    f.Score,
		Features: f.Features,
	}
}
