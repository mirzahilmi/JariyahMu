package main

import (
	"fmt"

	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(&viper)
	log := config.NewLogger(&viper)
	db := config.NewDatabase(&viper)
	validator := config.NewValidator()
	paseto := helper.NewPaseto()

	config.Bootstrap(&config.Config{
		Viper:    &viper,
		App:      app,
		DB:       &db,
		Log:      log,
		Validate: &validator,
		Paseto:   paseto,
	})

	serverPort := viper.GetInt("APP_PORT")

	err := app.Listen(fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
