package main

import (
	"context"
	"litetorrent-tracker/config"
	"litetorrent-tracker/internal"
	"litetorrent-tracker/internal/db"
	"litetorrent-tracker/pkg/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := db.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := internal.NewApp(cfg, dbClient).Init()
	srv := server.NewServer(cfg.Port, app.GetApiRouter())

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started at", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
