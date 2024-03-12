package rest

import "errors"

var (
	ErrMissingID    = errors.New("missing attempt id")
	ErrMissingToken = errors.New("missing attempt token")
)
