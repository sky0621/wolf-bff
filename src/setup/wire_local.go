//+build wireinject

package setup

import (
	"context"

	"github.com/google/wire"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (app, error) {
	wire.Build(
		connectDB,
		controller.NewResolver,
		setupRouter,
		newApp,
	)
	return app{}, nil
}
