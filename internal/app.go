package internal

import (
	"database/sql"
	"github.com/gorilla/mux"
	"litetorrent-tracker/config"
	"litetorrent-tracker/internal/repositories"
	"litetorrent-tracker/internal/transport"
)

type App struct {
	config   *config.Config
	dbClient *sql.DB

	handlers struct {
		peerHandler *transport.PeerHandler
	}
}

func NewApp(config *config.Config, db *sql.DB) *App {
	return &App{
		config:   config,
		dbClient: db,
	}
}

func (a *App) Init() *App {
	repo := repositories.NewRepo(a.dbClient)
	a.handlers.peerHandler = transport.NewHandler(repo)
	return a
}

func (a *App) GetApiRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/peers", a.handlers.peerHandler.CreatePeer).Methods("POST")
	r.HandleFunc("/peers/{peerId}", a.handlers.peerHandler.DeletePeer).Methods("DELETE")
	r.Path("/peers/{peerId}").HandlerFunc(a.handlers.peerHandler.UpdateDistributedFiles).Methods("PUT")
	r.Path("/peers").Queries("field", "{field}").HandlerFunc(a.handlers.peerHandler.GetPeersAddressesByFile).Methods("GET")
	return r
}
