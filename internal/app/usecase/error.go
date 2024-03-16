package usecase

import "errors"

var (
	ErrUserNotActive        = errors.New("user is not active")
	ErrVerificationNotExist = errors.New("verification attempt is not valid")
	ErrEmailExist           = errors.New("email already exist")
)
