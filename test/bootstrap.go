package test

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
)

var (
	viper     = config.NewViper()
	app       = config.NewFiber(&viper)
	log       = config.NewLogger(&viper)
	db        = config.NewDatabase(&viper)
	validator = config.NewValidator()
	pasetoo   = helper.NewPaseto()
)

func init() {
	config.Bootstrap(&config.Config{
		Viper:    &viper,
		App:      app,
		DB:       &db,
		Log:      log,
		Validate: &validator,
		Paseto:   pasetoo,
	})
}
