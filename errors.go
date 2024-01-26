package securedor

import "errors"

var (
	ErrNotImplemented      = errors.New("not implemented")
	ErrIncorrectAccessType = errors.New("invalid access type specified")
	ErrIncorrectActor      = errors.New("invalid actor specified")
	ErrAccessDenied        = errors.New("access denied")
)
