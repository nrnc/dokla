package fetch

import (
	"context"
	"errors"
	"time"

	"github.com/unbxd/go-base/utils/log"
)

type (
	Service interface {
		FetchById(ctx context.Context, req *Request) (*Response, error)
		FetchByDuration(ctx context.Context, req *Request) (*Response, error)
	}
	service struct {
		logger log.Logger
		repo   Repository
	}
)

func (s *service) FetchById(ctx context.Context, req *Request) (*Response, error) {
	s.logger.Info("fetching posts by id")
	r, err := s.repo.FetchById(ctx, req)
	if err != nil {
		s.logger.Error("fetchig failed " + err.Error())
		return nil, err
	}
	return r, nil
}

func formatAfterAndBefore(req *Request) error {

	if req.After == "" && req.Before == "" {
		req.After = time.Now().AddDate(0, 0, -3).Format(time.RFC3339)
		req.Before = time.Now().Format(time.RFC3339)
	} else if req.After == "" {

		t, err := time.Parse("02-01-2006", req.Before)
		if err != nil {
			return errors.New("invalid request")
		}
		req.Before = t.Format(time.RFC3339)
		req.After = t.AddDate(0, 0, -3).Format(time.RFC3339)

	} else if req.Before == "" {

		t, err := time.Parse("02-01-2006", req.After)
		if err != nil {
			return errors.New("invalid request")
		}

		req.After = t.Format(time.RFC3339)
		req.Before = t.AddDate(0, 0, 3).Format(time.RFC3339)

	} else {

		t, err := time.Parse("02-01-2006", req.After)
		if err != nil {
			return errors.New("invalid request")
		}
		req.After = t.Format(time.RFC3339)

		t, err = time.Parse("02-01-2006", req.Before)
		if err != nil {
			return errors.New("invalid request")
		}
		req.Before = t.Format(time.RFC3339)

	}

	return nil
}

func (s *service) FetchByDuration(ctx context.Context, req *Request) (*Response, error) {
	s.logger.Info("fetching posts by duration")

	err := formatAfterAndBefore(req)
	if err != nil {
		s.logger.Error("fetching failed " + err.Error())
		return nil, err
	}
	r, err := s.repo.FetchByDuration(ctx, req)

	if err != nil {
		s.logger.Error("fetching failed " + err.Error())
		return nil, err
	}

	return r, nil
}

// NewService returns default implementation of Service interface
func NewService(
	logger log.Logger,
	repository Repository,
) (Service, error) {
	a := service{
		logger: logger,
		repo:   repository,
	}

	return &a, nil
}
