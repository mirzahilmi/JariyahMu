package usecase

import "errors"

var (
	ErrUserNotActive        = errors.New("user is not active")
	ErrUserNotExist         = errors.New("user does not exist")
	ErrWrongPassword        = errors.New("password is incorrect")
	ErrVerificationNotExist = errors.New("verification attempt is not valid")
	ErrEmailExist           = errors.New("email already exist")
)
