package ingest

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
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
	source := ireq.Source
	tenant := ireq.Tenant

	r, err := m.client.Database(tenant).Collection(source+app).InsertOne(ctx, ireq)

	if err != nil {
		return nil, err
	}
	return r.InsertedID, err
}
