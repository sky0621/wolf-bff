package gateway

import (
	"context"

	"github.com/sky0621/wolf-bff/src/domain"
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
)

func NewWht(log system.Logger, db driver.RDB) domain.Wht {
	return &wht{log: log, db: db}
}

type wht struct {
	log system.Logger

	db driver.RDB
}

func (w *wht) FindWht(ctx context.Context) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("gateway.wht.FindWht__START")

	// FIXME:
}

func (w *wht) CreateWht(ctx context.Context) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("gateway.wht.CreateWht__START")

	// FIXME:
	//in := mdl.WHT{}
	//in.Insert()
}
