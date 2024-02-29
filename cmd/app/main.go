package main

import (
	"fmt"

	"github.com/MirzaHilmi/JariyahMu/internal/app/config"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(&viper)
	log := config.NewLogger(&viper)
	db := config.NewDatabase(&viper)
	validator := config.NewValidator()

	config.Bootstrap(&config.Config{
		Viper:    &viper,
		App:      app,
		DB:       &db,
		Log:      log,
		Validate: &validator,
	})

	serverPort := viper.GetInt("APP_PORT")

	err := app.Listen(fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
