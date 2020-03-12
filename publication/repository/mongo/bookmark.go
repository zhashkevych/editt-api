package mongo

import (
	"context"
	"edittapi/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Publication struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Author          string             `bson:"author"`
	Tags            []string           `bson:"tags"`
	Body            string             `bson:"body"`
	ImageLink       string             `bson:"imageLink"`
	Views           int32              `bson:"views"`
	Claps           int32              `bson:"claps"`
	ReadTimeMinutes int32              `bson:"readTime"`
	PublishedAt     time.Time          `bson:"publishedAt"`
}

type PublicationRepository struct {
	db *mongo.Collection
}

func NewPublicationRepository(db *mongo.Database, collection string) *PublicationRepository {
	return &PublicationRepository{
		db: db.Collection(collection),
	}
}

func (r PublicationRepository) Create(ctx context.Context, publication models.Publication) error {
	model := toPublication(publication)

	_, err := r.db.InsertOne(ctx, model)
	if err != nil {
		log.Errorf("Publication Repo: error occured on inserting publication: %s", err.Error())
		return err
	}

	return nil
}

func (r PublicationRepository) GetPopular(ctx context.Context, limit int64) ([]*models.Publication, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"views", -1}})
	opts.SetLimit(limit)

	var publications []*Publication

	cur, err := r.db.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
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

func (r PublicationRepository) GetLatest(ctx context.Context, limit int64) ([]*models.Publication, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})
	opts.SetLimit(limit)

	var publications []*Publication

	cur, err := r.db.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		log.Errorf("Publication Repo: error occured while finding latest publications: %s", err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem Publication
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("Publication Repo: error occured while decoding latest publications: %s", err.Error())
			return nil, err
		}

		publications = append(publications, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Errorf("Publication Repo: error occured while finding latest publications: %s", err.Error())
		return nil, err
	}

	if err := cur.Close(ctx); err != nil {
		log.Errorf("Publication Repo: error occured while finding latest publications: %s", err.Error())
		return nil, err
	}

	return toModels(publications), nil
}

func toPublication(p models.Publication) *Publication {
	return &Publication{
		Author:          p.Author,
		Tags:            p.Tags,
		Body:            p.Body,
		ImageLink:       p.ImageLink,
		Views:           p.Views,
		Claps:           p.Claps,
		ReadTimeMinutes: p.ReadTimeMinutes,
		PublishedAt:     p.PublishedAt,
	}
}

func toModel(p *Publication) *models.Publication {
	return &models.Publication{
		Author:          p.Author,
		Tags:            p.Tags,
		Body:            p.Body,
		ImageLink:       p.ImageLink,
		Views:           p.Views,
		Claps:           p.Claps,
		ReadTimeMinutes: p.ReadTimeMinutes,
		PublishedAt:     p.PublishedAt,
	}
}

func toModels(ps []*Publication) []*models.Publication {
	out := make([]*models.Publication, len(ps))
	for i := range ps {
		out[i] = toModel(ps[i])
	}

	return out
}
