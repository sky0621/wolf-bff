//+build wireinject

package setup

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (app, error) {
	wire.Build(
		connectLocalDB,
		controller.NewResolver,
		setupRouter,
		newApp,
	)
	return app{}, nil
}

func connectLocalDB(cfg config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSSLMode)

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
