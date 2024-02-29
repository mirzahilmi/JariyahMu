package usecase

import (
	"context"

	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
)

type UserUsecaseItf interface {
	RegisterUser(ctx context.Context, user model.StoreUser) (string, error)
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

func (usc *UserUsecase) RegisterUser(ctx context.Context, user model.StoreUser) (string, error) {
	id, err := helper.NewULID()
	if err != nil {
		return "", nil
	}

	user.ID = id.String()
	if err := usc.repo.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := usc.encode(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
