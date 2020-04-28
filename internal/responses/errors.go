package responses

import "errors"

var (
	ErrPageNorFound = errors.New("not found")
	ErrInvalidData  = errors.New("invalid data")
)
