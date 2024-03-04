package usecase

import (
	"context"

	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
)

type UserUsecaseItf interface {
	RegisterUser(ctx context.Context, user model.CreateUserRequest) (string, error)
}

type UserUsecase struct {
	repo   repository.UserRepositoryItf
	encode func(id string) (string, error)
}

func NewUserUsecase(
	repo repository.UserRepositoryItf,
	encode func(id string) (string, error),
) UserUsecaseItf {
	return &UserUsecase{repo: repo, encode: encode}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, userRequest model.CreateUserRequest) (string, error) {
	id, err := helper.ULID()
	if err != nil {
		return "", nil
	}

	hashed, err := helper.BcryptHash(userRequest.Password)
	if err != nil {
		return "", nil
	}

	user := model.StoreUser{
		ID:             id,
		FullName:       userRequest.FullName,
		Email:          userRequest.Email,
		HashedPassword: hashed,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := u.encode(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
