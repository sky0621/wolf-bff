package driver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/system"
	"golang.org/x/xerrors"
)

type Web interface {
	Start() error
}

type web struct {
	cfg system.Config
	log system.Logger

	r *chi.Mux
}

func NewWeb(cfg system.Config, log system.Logger, resolver controller.ResolverRoot) Web {
	r := chi.NewRouter()

	// FIXME: 本番はNG
	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", handler.NewDefaultServer(controller.NewExecutableSchema(controller.Config{Resolvers: resolver})))

	return &web{cfg: cfg, log: log, r: r}
}

func (w *web) Start() error {
	if w == nil {
		return xerrors.New("web is nil")
	}
	if err := http.ListenAndServe(":"+w.cfg.GetWebSetting().WebPort, w.r); err != nil {
		return xerrors.Errorf("failed to ListenAndServe: %w", err)
	}
	return nil
}
