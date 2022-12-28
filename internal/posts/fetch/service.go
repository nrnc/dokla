package fetch

import (
	"context"

	"github.com/unbxd/go-base/utils/log"
)

type (
	Service interface {
		FetchPosts(ctx context.Context, req *Request) (*Response, error)
	}
	service struct {
		logger log.Logger
	}
)

func (s *service) FetchPosts(ctx context.Context, req *Request) (*Response, error) {
	r := Response{}
	s.logger.Info("fetching the posts")
	return &r, nil
}

// NewService returns default implementation of Service interface
func NewService(
	logger log.Logger,
) (Service, error) {
	a := service{
		logger: logger,
	}

	return &a, nil
}
