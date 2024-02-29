package model

import "time"

type StoreResetAttempt struct {
	UserID     string
	Token      string
	ValidUntil time.Time
}
