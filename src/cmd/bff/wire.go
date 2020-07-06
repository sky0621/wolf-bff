//+build wireinject

package main

import (
	"context"

	"github.com/sky0621/wolf-bff/src/adapter/gateway"
	"github.com/sky0621/wolf-bff/src/usecase"

	"github.com/google/wire"
	wht "github.com/sky0621/wolf-bff/src"
	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
)

var commonSet = wire.NewSet(
	/*
	 * adapter/gateway
	 */
	gateway.NewWht,

	/*
	 * usecase
	 */
	usecase.NewWht,

	/*
	 * adapter/controller
	 */
	controller.NewResolver,

	// Webサーバー
	driver.NewWeb,

	// アプリケーションそのもの
	wht.NewApp,
)

// デプロイ先インフラで動かす際の設定
func build(ctx context.Context, cfg system.Config) (wht.App, error) {
	wire.Build(
		// ロガー
		system.NewLogger,

		// RDBコネクション
		driver.NewRDB,

		// ユースケースやドメインロジック、アダプター等
		commonSet,
	)
	return nil, nil
}
