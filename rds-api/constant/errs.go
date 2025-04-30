package constant

import "errors"

const (
	CodeSuccess           = 200
	CodeErrBadRequest     = 400
	CodeErrUnauthorized   = 401
	CodeErrInternalServer = 500
)

var (
	ErrTypeInternalServer = errors.New("ErrTypeInternalServer")
	ErrTypeInvalidParams  = errors.New("ErrTypeInvalidParams")
)
