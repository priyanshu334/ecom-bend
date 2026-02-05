package auth

import "errors"

var (
	ErrorUserAlreadyExists  = errors.New("user already exists")
	ErrorInvalidCredintials = errors.New("invalid credintials")
)
