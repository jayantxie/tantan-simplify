package errors

import "errors"

var (
	ErrUnknown      = errors.New("unknown internal error")
	ErrOperated     = errors.New("operated")
	ErrUnknownState = errors.New("unknown state")
	ErrUserNotFound = errors.New("user not found")
)
