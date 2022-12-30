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
	r := Response{}
	s.logger.Info("fetching the posts")
	ir, err := s.repo.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	r.Id = ir
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
