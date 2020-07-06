package usecase

import (
	"context"

	"github.com/sky0621/wolf-bff/src/domain"
	"github.com/sky0621/wolf-bff/src/system"
)

type Wht interface {
	FindWht(ctx context.Context)
	CreateWht(ctx context.Context)
}

type wht struct {
	log system.Logger

	whtLogic domain.Wht
}

func NewWht(log system.Logger, whtLogic domain.Wht) Wht {
	return &wht{log: log, whtLogic: whtLogic}
}

func (w *wht) FindWht(ctx context.Context) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("usecase.Wht.FindWht__START")
	w.whtLogic.FindWht(ctx)
}

func (w *wht) CreateWht(ctx context.Context) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("usecase.Wht.CreateWht__START")
	w.whtLogic.CreateWht(ctx)
}
