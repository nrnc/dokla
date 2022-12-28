package ingest

import "context"

type Service interface {
	Ingest(ctx context.Context, req Request) Response
}
