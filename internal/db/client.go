package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"litetorrent-tracker/config"
)

func NewClient(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error while opening db connection: %w", err)
	}
	return db, nil
}
