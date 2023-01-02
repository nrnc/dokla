package ingest

import (
	"context"

	"github.com/unbxd/go-base/utils/log"
)

type (
	Service interface {
		Ingest(ctx context.Context, req *Request) (*Response, error)
	}
	service struct {
		logger log.Logger
		repo   Repository
	}
)

func (s *service) Ingest(ctx context.Context, req *Request) (*Response, error) {

	s.logger.Info("ingesting the posts")
	ir, err := s.repo.InsertOne(ctx, req)
	if err != nil {
		s.logger.Error("failed with the following error" + err.Error())
		return nil, err
	}

	r := Response{Id: ir, Success: true}

	return &r, nil
}

// NewService returns default implementation of Service interface
func NewService(
	logger log.Logger,
	repo Repository,
) (Service, error) {
	a := service{
		logger: logger,
		repo:   repo,
	}

	return &a, nil
}
