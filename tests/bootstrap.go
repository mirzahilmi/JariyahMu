package tests

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
)

var (
	viper     = config.NewViper()
	app       = config.NewFiber(&viper)
	log       = config.NewLogger(&viper)
	db        = config.NewDatabase(&viper)
	validator = config.NewValidator()
)

func init() {
	config.Bootstrap(&config.Config{
		Viper:    &viper,
		App:      app,
		DB:       &db,
		Log:      log,
		Validate: &validator,
	})
}
