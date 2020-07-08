package usecase

import (
	"context"

	"github.com/sky0621/wolf-bff/src/domain"
	"github.com/sky0621/wolf-bff/src/system"
)

type Wht struct {
}

type WhtInputPort interface {
	FindWht(ctx context.Context) ([]*Wht, error)
	CreateWht(ctx context.Context)
}

type WhtOutputPort interface {
	Output()
}

type whtInteractor struct {
	log      system.Logger
	whtLogic domain.WhtLogic
}

func NewWht(log system.Logger, whtLogic domain.WhtLogic) WhtInputPort {
	return &whtInteractor{log: log, whtLogic: whtLogic}
}

func (w *whtInteractor) FindWht(ctx context.Context) ([]*Wht, error) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("usecase.WhtInputPort.FindWht__START")
	w.whtLogic.FindWht(ctx)
}

func (w *whtInteractor) CreateWht(ctx context.Context) {
	l := w.log.WithRequestContext(ctx)
	l.Info().Msg("usecase.WhtInputPort.CreateWht__START")
	w.whtLogic.CreateWht(ctx)
}
