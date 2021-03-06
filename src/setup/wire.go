//+build wireinject

package setup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
	dsn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s:asia-northeast1:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.ProjectID, cfg.DBHost, cfg.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	// FIXME: 本番はNG?
	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	boil.SetLocation(loc)

	return db, nil
}

func setupRouter(cfg config, resolver *controller.Resolver) *chi.Mux {
	r := chi.NewRouter()
	// FIXME: 本番はNG
	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", controller.DataLoaderMiddleware(resolver, graphQlServer(resolver)))
	return r
}

func graphQlServer(resolver *controller.Resolver) *handler.Server {
	cfg := controller.Config{Resolvers: resolver}
	// FIXME: 認可実装
	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role gqlmodel.Role) (interface{}, error) {
		// or let it pass through
		return next(ctx)
	}
	srv := handler.New(controller.NewExecutableSchema(cfg))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     128 * mb,
		MaxUploadSize: 100 * mb,
	})

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Println(err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Println(err)
		return fmt.Errorf("internal server error! %v", err)
	})

	return srv
}
