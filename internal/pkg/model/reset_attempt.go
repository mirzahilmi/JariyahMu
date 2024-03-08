package model

import "time"

type StoreResetAttempt struct {
	ID         string    `db:"ID"`
	UserID     string    `db:"UserID"`
	Token      string    `db:"Token"`
	Expiration time.Time `db:"Expiration"`
}
