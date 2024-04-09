package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port             string
	ConnectionString Server
}

type Server struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func DBConnection(cfg Config) (*sqlx.DB, error) {
	db := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=%s",
		cfg.ConnectionString.Server, cfg.ConnectionString.Port, cfg.ConnectionString.User, cfg.ConnectionString.Password, cfg.ConnectionString.Database, cfg.ConnectionString.SSLMode)

	conn, err := sqlx.Open("postgres", db)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to db")

	return conn, nil
}
