//+build wireinject

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/adapter/controller/graphqlmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

// デプロイ先インフラで動かす際の設定
func build(ctx context.Context, cfg config) (app, error) {
	wire.Build(
		connectDB,
		controller.NewResolver,
		setupRouter,
		newApp,
	)
	return app{}, nil
}

func connectDB(cfg config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	boil.SetLocation(loc)

	return db, nil
}

func setupRouter(cfg config, resolver controller.ResolverRoot) *chi.Mux {
	r := chi.NewRouter()

	// FIXME: 本番はNG
	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))

	c := controller.Config{Resolvers: resolver}
	// FIXME: 認可実装
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role graphqlmodel.Role) (interface{}, error) {
		// or let it pass through
		return next(ctx)
	}
	r.Handle("/query", handler.NewDefaultServer(controller.NewExecutableSchema(c)))
	return r
}
