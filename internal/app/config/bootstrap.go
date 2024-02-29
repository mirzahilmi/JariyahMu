package config

import (
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

func Bootstrap(config *Config) {
	encode, _ := NewPaseto(config.Viper)

	userRepository := repository.NewUserRepository(config.DB)
	_ = usecase.NewUserUsecase(userRepository, encode)

	api := config.App.Group("/api/v1")

	rest.RegisterUtilsHandler(api)
}
