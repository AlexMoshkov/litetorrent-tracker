package internal

import (
	"database/sql"
	"litetorrent-tracker/config"
	"litetorrent-tracker/internal/transport"
)

type App struct {
	Config   *config.Config
	DBClient *sql.DB
	Router   *transport.Router
}

func NewApp(config *config.Config, db *sql.DB) *App {
	return &App{
		Config:   config,
		DBClient: db,
		Router:   transport.NewRouter(),
	}
}

func (a *App) Init() *App {
	// ...
	return a
}
