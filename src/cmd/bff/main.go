package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	wht "github.com/sky0621/wolf-bff/src"
	"github.com/sky0621/wolf-bff/src/system"
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
	cfg := system.NewConfig()
	if cfg == nil {
		log.Println("config is nil")
		return abnormalEnd
	}

	ctx := context.Background()
	var app wht.App
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
		if app != nil {
			if err := app.Shutdown(); err != nil {
				log.Printf("%+v", err)
			}
		}
	}()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		if err := app.Shutdown(); err != nil {
			log.Printf("%+v", err)
		}
		os.Exit(abnormalEnd)
	}()

	if err := app.Start(); err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	return normalEnd
}
