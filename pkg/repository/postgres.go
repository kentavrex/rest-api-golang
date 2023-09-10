package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable         = "users"
	segmentsTable      = "segments"
	usersSegmentsTable = "users_segments"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("sslmode=disable host=%s user=%s dbname=%s password=%s",
			cfg.Host, cfg.Username, cfg.DBName, cfg.Password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))
		return nil, err
	}

	return db, nil
}
