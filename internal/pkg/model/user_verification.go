package model

type UserVerificationResource struct {
	ID     string `db:"ID"`
	UserID string `db:"UserID"`
	Token  string `db:"Token"`
}
