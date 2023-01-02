package fetch

import (
	"context"
	"time"

	"github.com/nrnc/dokla/cmd/dokla/flags"
	"github.com/unbxd/go-base/utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository interface {
	FetchById(ctx context.Context, request *Request) (*Response, error)
	FetchByDuration(ctx context.Context, request *Request) (*Response, error)
}

type mongoRepo struct {
	client *mongo.Client
	logger log.Logger
}

func NewMongoFetchRepo(client *mongo.Client, logger log.Logger) *mongoRepo {
	return &mongoRepo{
		logger: logger,
		client: client,
	}
}

func (m *mongoRepo) FetchById(ctx context.Context, req *Request) (*Response, error) {

	tenant := req.Tenant
	app := req.App
	id := req.Id

	filter := bson.D{{"post_id", bson.D{{"$eq", id}}}}

	opts := options.Collection().SetReadPreference(readpref.Secondary())

	ctx, cancel := context.WithTimeout(ctx, time.Duration(flags.DbReadTimeout)*time.Millisecond)
	defer cancel()

	cursor, err := m.client.Database(tenant).Collection(app, opts).Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	var posts []Post
	if err = cursor.All(ctx, &posts); err != nil {
		m.logger.Error("fetching failed with error : " + err.Error())
		return nil, err
	}

	cursor.Close(ctx)

	return &Response{Posts: posts}, err
}

func (m *mongoRepo) FetchByDuration(ctx context.Context, req *Request) (*Response, error) {

	tenant := req.Tenant
	app := req.App
	after := req.After
	before := req.Before

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"created_at", bson.D{{"$gte", after}}}},
				bson.D{{"created_at", bson.D{{"$lte", before}}}},
			},
		},
	}

	opts := options.Collection().SetReadPreference(readpref.Secondary())

	ctx, cancel := context.WithTimeout(ctx, time.Duration(flags.DbReadTimeout)*time.Millisecond)
	defer cancel()

	cursor, err := m.client.Database(tenant).Collection(app, opts).Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	var posts []Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return &Response{Posts: posts}, err
}
