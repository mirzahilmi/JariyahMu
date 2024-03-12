package model

type UserResource struct {
	ID             string `db:"ID"`
	FullName       string `db:"FullName"`
	Email          string `db:"Email"`
	HashedPassword string `db:"HashedPassword"`
	ProfilePicture string `db:"ProfilePicture"`
	Active         bool   `db:"Active"`
}

type CreateUserRequest struct {
	FullName             string `json:"fullName" validate:"required,ascii"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

type UserLoginAttemptRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type QueryUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type AttemptResetPasswordRequest struct {
	ID    string `json:"id" validate:"required,ulid"`
	Token int    `json:"token" validate:"required,len=6"`
}
