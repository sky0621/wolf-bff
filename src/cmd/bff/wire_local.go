//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	wht "github.com/sky0621/wolf-bff/src"
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg system.Config) (wht.App, error) {
	wire.Build(
		// ロガー
		system.NewLogger,

		// RDBコネクション
		driver.NewRDB,

		// Webサーバー
		driver.NewWeb,

		// アプリケーションそのもの
		wht.NewApp,

		// ユースケースやドメインロジック、アダプター
		//domainSet,
	)
	return nil, nil
}
