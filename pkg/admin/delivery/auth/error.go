package auth

import "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrInvalidCredentials = errors.New("invalid credentials")
