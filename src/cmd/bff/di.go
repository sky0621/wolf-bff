//+build wireinject

package main

import (
	"github.com/google/wire"
	wht "github.com/sky0621/wolf-bff/src"
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
)

var set = wire.NewSet(
	// ロガー
	system.NewLogger,

	// Config => RDBコネクションプール
	driver.NewRDB,

	// RDBコネクションプール => domain層インタフェースをadapter層で実装
	//gateway.NewItem,
	//gateway.NewItemHolder,

	// domain層インタフェース => usecase層
	//usecase.NewItem,
	//usecase.NewItemHolder,

	// usecase層 => GraphQLリゾルバー
	//controller.NewResolverRoot,

	// WebFramework
	driver.NewWeb,

	// システムそのもの
	wht.NewApp,
)

func di(cfg system.Config) wht.App {
	wire.Build(set)
	return nil
}
