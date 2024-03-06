package config

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/delivery/rest"
	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Viper    *viper.Viper
	App      *fiber.App
	DB       *sqlx.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Paseto   helper.Paseto
}

func Bootstrap(conf *Config) {
	router := conf.App.Group("/api/v1")
	rest.RegisterUtilsHandler(router)

	userRepository := repository.NewUserRepository(conf.DB)
	userUsecase := usecase.NewUserUsecase(userRepository, conf.Paseto, conf.Viper)
	rest.RegisterUserHandler(userUsecase, router)
}
