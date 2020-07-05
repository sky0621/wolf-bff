package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
		log.Fatal()
	}

	//ctx := context.Background()

	app := di(cfg)
	defer app.Shutdown()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		app.Shutdown()
		os.Exit(int(abnormalEnd))
	}()

	if err := app.Start(); err != nil {
		fmt.Errorf("the application failed to start")
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
