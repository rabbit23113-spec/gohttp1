package database

import (
	"database/sql"
	"fmt"
	"log"
	"main/package/cfg"
)

func NewPostgresDb(conf *cfg.Config) (*sql.DB, error) {
	dsn := conf.DatabaseUrl()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Database connection error: %e", err)
	}
	fmt.Println("Successfully connected to the database!")
	return db, nil
}
