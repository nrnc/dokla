package utils

import (
	"context"
	"time"

	"github.com/nrnc/dokla/cmd/dokla/flags"
	"github.com/pkg/errors"
	"github.com/unbxd/go-base/utils/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoClient(conn string, logger log.Logger) (*mongo.Client, error) {
	var (
		cl  *mongo.Client
		cx  = context.Background()
		err error
	)

	credential := options.Credential{
		Username: flags.DbRootUser,
		Password: flags.DbRootPassword,
	}

	logger.Info("conn details", log.String("conn", conn))

	clientOpts := options.Client().ApplyURI(conn).
		SetConnectTimeout(time.Duration(flags.DbConnTimeout) * time.Millisecond).
		SetAuth(credential)

	cl, err = mongo.Connect(cx, clientOpts)

	if err != nil {
		logger.Error("connecting to mongodb", log.String("err", err.Error()))
		return nil, errors.Wrap(
			err, "failed to connect to mongo",
		)
	}

	err = cl.Ping(cx, readpref.Primary())
	if err != nil {
		logger.Error("connecting to mongodb", log.String("err", err.Error()))
		return nil, errors.Wrap(
			err, "failed to ping mongo server",
		)
	}

	return cl, err
}
