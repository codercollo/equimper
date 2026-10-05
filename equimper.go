package equimper

import "errors"

var (
	ErrValidation     = errors.New("validation error")
	ErrNotFound       = errors.New("not found")
	ErrBadCredentials = errors.New("email or password wrong combination")
)
