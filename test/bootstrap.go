package test

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/auth"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/email"
)

var (
	viperInstance     = config.NewViper()
	app               = config.NewFiber(&viperInstance)
	log               = config.NewLogger(&viperInstance)
	db                = config.NewDatabase(&viperInstance)
	validatorInstance = config.NewValidator()
	pasetoo           = auth.NewPaseto()
	mailer            = email.MailerMock{}
)

func init() {
	config.Bootstrap(&config.Config{
		Viper:    &viperInstance,
		App:      app,
		DB:       &db,
		Log:      log,
		Validate: &validatorInstance,
		Mailer:   &mailer,
		Paseto:   &pasetoo,
	})
}
