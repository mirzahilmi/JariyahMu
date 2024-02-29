package model

type StoreUser struct {
	ID             string
	FullName       string
	Email          string
	HashedPassword string
	ProfilePicture string
}

type CreateUserRequest struct {
	FullName             string `json:"fullName" validate:"required,ascii"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}
