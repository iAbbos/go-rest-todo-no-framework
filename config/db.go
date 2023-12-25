package config

import (
	"database/sql"
	"fmt"
	"golang-crud/helpers"

	_ "github.com/lib/pq" // Postgres goalng driver
	"github.com/rs/zerolog/log"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "todo_db"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", sqlInfo)
	helpers.PanicIfError(err)

	err = db.Ping()
	helpers.PanicIfError(err)

	log.Info().Msg("Database connection established")

	return db
}
