package usecase

import (
	"context"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/spf13/viper"
)

type UserUsecaseItf interface {
	RegisterUser(ctx context.Context, user model.CreateUserRequest) (string, error)
}

type UserUsecase struct {
	repo           repository.UserRepositoryItf
	pasetoInstance helper.Paseto
	viper          *viper.Viper
}

func NewUserUsecase(
	repo repository.UserRepositoryItf,
	pasetoInstance helper.Paseto,
	viper *viper.Viper,
) UserUsecaseItf {
	return &UserUsecase{repo, pasetoInstance, viper}
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

	token := paseto.NewToken()

	token.SetAudience("*")
	token.SetIssuer(viper.GetString("APP_HOST"))
	token.SetSubject(id)
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetNotBefore(time.Now())
	token.SetIssuedAt(time.Now())

	signed, err := u.pasetoInstance.Encode(token)
	if err != nil {
		return "", err
	}

	return signed, nil
}
