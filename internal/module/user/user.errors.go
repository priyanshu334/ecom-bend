package user

import "errors"

var (
	ErrProfileNotFound = errors.New("user profile not found")
	ErrAddressNotFound = errors.New("address not found")
	ErrUnautorized     = errors.New("ErrUnautorized access")
)
