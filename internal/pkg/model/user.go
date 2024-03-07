package model

type StoreUser struct {
	ID             string `db:"ID"`
	FullName       string `db:"FullName"`
	Email          string `db:"Email"`
	HashedPassword string `db:"HashedPassword"`
	ProfilePicture string `db:"ProfilePicture"`
}

type CreateUserRequest struct {
	FullName             string `json:"fullName" validate:"required,ascii"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}
