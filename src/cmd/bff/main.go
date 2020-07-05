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
		log.Println(err)
		return abnormalEnd
	}
	defer func() {
		if app != nil {
			app.Shutdown()
		}
	}()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		app.Shutdown()
		os.Exit(abnormalEnd)
	}()

	if err := app.Start(); err != nil {
		log.Println(err)
		return abnormalEnd
	}
	return normalEnd

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)
	//
	//port := cfg.GetWebSetting().WebPort
	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}
