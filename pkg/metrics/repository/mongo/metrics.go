package mongo

import (
	"context"
	"edittapi/pkg/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Metrics struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	UniqueVisitorsCount int64              `bson:"unique_visitors_count"`
	Timestamp           time.Time              `bson:"timestamp"`
}

type MetricsRepository struct {
	db *mongo.Collection
}

func NewMetricsRepository(db *mongo.Database, collection string) *MetricsRepository {
	return &MetricsRepository{
		db: db.Collection(collection),
	}
}

func (r MetricsRepository) SetMetrics(ctx context.Context, metrics models.Metrics) error {
	model := toMetrics(metrics)

	_, err := r.db.InsertOne(ctx, model)
	if err != nil {
		log.Errorf("Publication Repo: error occured on inserting publication: %s", err.Error())
		return err
	}

	return nil
}

func (r MetricsRepository) GetMetrics(ctx context.Context, timeFrom time.Time) ([]*models.Metrics, error) {
	var ms []*Metrics

	opts := options.Find()
	opts.SetSort(bson.M{"_id": -1})

	//cur, err := r.db.Find(ctx, bson.M{"timestamp": bson.M{"$gte": timeFrom.Format(time.RFC3339)}})
	cur, err := r.db.Find(ctx, bson.M{"timestamp": bson.M{"$gte": timeFrom}}, opts)
	if err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	for cur.Next(ctx) {
		var elem Metrics
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("Publication Repo: error occured while decoding popular publications: %s", err.Error())
			return nil, err
		}

		ms = append(ms, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	if err := cur.Close(ctx); err != nil {
		log.Errorf("Publication Repo: error occured while finding popular publications: %s", err.Error())
		return nil, err
	}

	return toModels(ms), nil
}

func toMetrics(m models.Metrics) *Metrics {
	return &Metrics{
		UniqueVisitorsCount: m.UniqueVisitorsCount,
		Timestamp:           m.Timestamp,
	}
}

func toModel(m *Metrics) *models.Metrics {
	return &models.Metrics{
		UniqueVisitorsCount: m.UniqueVisitorsCount,
		Timestamp:           m.Timestamp,
	}
}

func toModels(ms []*Metrics) []*models.Metrics {
	out := make([]*models.Metrics, len(ms))
	for i := range ms {
		out[i] = toModel(ms[i])
	}

	return out
}
