package usecase

import "errors"

var (
	ErrUserNotActive                  = errors.New("user is not active")
	ErrVerificationNotExist           = errors.New("verification attempt is not valid")
	ErrFailToUpdateVerificationStatus = errors.New("fail to update verification status")
)
