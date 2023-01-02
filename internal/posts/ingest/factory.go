package ingest

import (
	"errors"

	"github.com/nrnc/dokla/internal/posts/ingest/consts"
)

func ReqBuilderFactory(source string) (ReqBuilder, error) {
	switch source {
	case consts.PLAYSTORE:
		return &PlayStore{}, nil
	case consts.TWITTER:
		return &Twitter{}, nil
	case consts.DISCOURSE:
		return &Discourse{}, nil
	case consts.DEFAULT:
		return &Request{}, nil
	default:
		return nil, errors.New("source specified is not supported. we might be adding in future")
	}
}
