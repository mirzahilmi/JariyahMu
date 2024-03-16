package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewDatabase(viper *viper.Viper) sqlx.DB {
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	schema := viper.GetString("DB_SCHEMA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("fatal error database: %w", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("fatal error database connection: %w", err))
	}

	return *db
}
