package ingest

import (
	"context"
	"time"

	"github.com/nrnc/dokla/cmd/dokla/flags"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertOne(ctx context.Context, request interface{}) (interface{}, error)
}

type mongoRepo struct {
	client *mongo.Client
}

func NewMongoIngestRepo(client *mongo.Client) *mongoRepo {
	return &mongoRepo{
		client: client,
	}
}

func (m *mongoRepo) InsertOne(ctx context.Context, request interface{}) (interface{}, error) {

	ireq := request.(*Request)
	app := ireq.App
	tenant := ireq.Tenant

	opts := options.Update().SetUpsert(true)

	filter := bson.M{"post_id": ireq.PostId}
	updatePost := bson.M{"$set": ireq}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(flags.DbWriteTimeout)*time.Millisecond)
	defer cancel()

	r, err := m.client.Database(tenant).Collection(app).UpdateOne(ctx, filter, updatePost, opts)

	if err != nil {
		return nil, err
	}
	return r.UpsertedID, err
}
