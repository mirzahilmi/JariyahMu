package test

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
	"github.com/MirzaHilmi/JariyahMu/internal/app/delivery/rest"
	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
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
}

func Bootstrap(conf *Config) {
	encode, _ := config.NewPaseto(conf.Viper)
	router := conf.App.Group("/api/v1")
	rest.RegisterUtilsHandler(router)

	userRepository := repository.NewUserRepository(conf.DB)
	userUsecase := usecase.NewUserUsecase(userRepository, encode)
	rest.RegisterUserHandler(userUsecase, router)
}
