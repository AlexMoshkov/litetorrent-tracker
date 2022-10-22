package internal

import (
	"litetorrent-tracker/config"
	"litetorrent-tracker/internal/transport"
)

type App struct {
	Config *config.Config
	Router *transport.Router
}

func NewApp(config *config.Config) *App {
	return &App{
		Config: config,
		Router: transport.NewRouter(),
	}
}

func (a *App) Init() *App {
	// ...
	return a
}
