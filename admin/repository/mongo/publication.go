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

type Publication struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Author      string             `bson:"author"`
	Title       string             `bson:"title"`
	Tags        []string           `bson:"tags"`
	Body        string             `bson:"body"`
	ImageLink   string             `bson:"imageLink"`
	Views       int32              `bson:"views"`
	Reactions   int32              `bson:"reactions"`
	ReadingTime int32              `bson:"readingTime"`
	PublishedAt time.Time          `bson:"publishedAt"`
}

type PublicationRepository struct {
	db *mongo.Collection
}

func NewPublicationRepository(db *mongo.Database, collection string) *PublicationRepository {
	return &PublicationRepository{
		db: db.Collection(collection),
	}
}

func (r PublicationRepository) GetPublications(ctx context.Context) ([]*models.Publication, error) {
	var publications []*Publication

	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	for cur.Next(ctx) {
		var elem Publication
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("Publication Repo: error occured while decoding popular publications: %s", err.Error())
			return nil, err
		}

		publications = append(publications, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	if err := cur.Close(ctx); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	return toModels(publications), nil
}

func (r PublicationRepository) RemovePublication(ctx context.Context, id string) error {
	pid, _ := primitive.ObjectIDFromHex(id)
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": pid})

	return err
}

func toPublication(p models.Publication) *Publication {
	return &Publication{
		Author:      p.Author,
		Title:       p.Title,
		Tags:        p.Tags,
		Body:        p.Body,
		ImageLink:   p.ImageLink,
		Views:       p.Views,
		Reactions:   p.Reactions,
		ReadingTime: p.ReadingTime,
		PublishedAt: p.PublishedAt,
	}
}

func toModel(p *Publication) *models.Publication {
	return &models.Publication{
		ID:          p.ID.Hex(),
		Author:      p.Author,
		Title:       p.Title,
		Tags:        p.Tags,
		Body:        p.Body,
		ImageLink:   p.ImageLink,
		Views:       p.Views,
		Reactions:   p.Reactions,
		ReadingTime: p.ReadingTime,
		PublishedAt: p.PublishedAt,
	}
}

func toModels(ps []*Publication) []*models.Publication {
	out := make([]*models.Publication, len(ps))
	for i := range ps {
		out[i] = toModel(ps[i])
	}

	return out
}
