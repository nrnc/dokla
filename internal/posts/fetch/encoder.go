package fetch

import (
	"context"
	"encoding/json"
	"net/http"
	nethttp "net/http"

	"github.com/nrnc/dokla/internal/posts/errors"
)

// encoder encodes the response
func encoder(ctx context.Context, w http.ResponseWriter, res interface{}) error {

	r := res.(*Response)

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(r)
}

// errorEncoder adds the error to the body and sends status 500
func errorEncoder(_ context.Context, err error, w nethttp.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")

	// always set to internal server error
	code := nethttp.StatusInternalServerError
	message := errors.DEFAULT_ERROR
	w.WriteHeader(code)

	errRes := &Response{
		Error: &Error{message, int32(code)},
	}
	json.NewEncoder(w).Encode(errRes)

}
