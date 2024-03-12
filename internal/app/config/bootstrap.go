package config

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/delivery/rest"
	"github.com/MirzaHilmi/JariyahMu/internal/app/repository"
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/auth"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/email"
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
	Paseto   *auth.Paseto
	Mailer   *email.VerificationMailer
}

func Bootstrap(conf *Config) {
	router := conf.App.Group("/api/v1")
	rest.RegisterUtilsHandler(router)

	userRepo := repository.NewUserRepository(conf.DB)
	userVerificationRepo := repository.NewUserVerificationRepository(conf.DB)
	resetAttemptRepo := repository.NewResetAttemptRepository(conf.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, userVerificationRepo, resetAttemptRepo, *conf.Mailer, conf.Paseto, conf.Viper)
	rest.RegisterUserHandler(userUsecase, router)
}
