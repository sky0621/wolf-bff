package main

import (
	"log"
	"net/http"

	"github.com/sky0621/wolf-bff/src/system"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sky0621/wolf-bff/src/graph"
	"github.com/sky0621/wolf-bff/src/graph/generated"
)

func main() {
	cfg := system.NewConfig()
	if cfg == nil {
		log.Fatal()
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := cfg.GetWebSetting().WebPort
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
