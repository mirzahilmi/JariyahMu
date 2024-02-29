package rest

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
)

type UserDelivery struct {
	usc usecase.UserUsecaseItf
}

func NewUserDelivery(usc usecase.UserUsecaseItf) UserDelivery {
	return UserDelivery{usc}
}
