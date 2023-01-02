package ingest

import (
	"errors"
	"net/http"
)

type ReqDecoder interface {
	Decode(req *http.Request) error
}

type Adaptor interface {
	Adapt() *Request
}

type PathParamsSetter interface {
	SetPathParams(pathParams *PathParams)
}

type TimeSetter interface {
	SetTime()
}
type Validator interface {
	IsValid() bool
}

type ReqBuilder interface {
	PathParamsSetter
	ReqDecoder
	Validator
	TimeSetter
	Adaptor
}

func BuildRequest(req *http.Request, pp *PathParams) (*Request, error) {

	request, err := ReqBuilderFactory(pp.Source)
	if err != nil {
		return nil, err
	}
	request.SetPathParams(pp)
	err = request.Decode(req)

	if err != nil {
		return nil, err
	}

	if !request.IsValid() {
		return nil, errors.New("not a valid request")
	}

	request.SetTime()
	br := request.Adapt()

	return br, nil
}
