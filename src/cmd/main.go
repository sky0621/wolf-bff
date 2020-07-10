package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type exitCode int

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(int(execMain()))
}

func execMain() exitCode {
	cfg := newConfig()
	ctx := context.Background()
	var app app
	var err error
	if cfg.IsLocal() {
		app, err = buildLocal(ctx, cfg)
	} else {
		app, err = build(ctx, cfg)
	}
	if err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	defer func() {
		if err := app.db.Close(); err != nil {
			log.Printf("%+v", err)
		}
	}()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		if err := app.db.Close(); err != nil {
			log.Printf("%+v", err)
		}
		os.Exit(abnormalEnd)
	}()

	if err := http.ListenAndServe(":"+cfg.WebPort, app.r); err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	return normalEnd
}

type app struct {
	db *sqlx.DB
	r  *chi.Mux
}

func newApp(db *sqlx.DB, r *chi.Mux) app {
	return app{
		db: db,
		r:  r,
	}
}
