package usecase

import (
	"context"

	"github.com/sky0621/wolf-bff/src/domain"
)

type Wht interface {
	FindWht(ctx context.Context)
	CreateWht(ctx context.Context)
}

type wht struct {
	whtLogic domain.Wht
}

func NewWht(whtLogic domain.Wht) Wht {
	return &wht{whtLogic: whtLogic}
}

func (w *wht) FindWht(ctx context.Context) {
	w.whtLogic.FindWht(ctx)
}

func (w *wht) CreateWht(ctx context.Context) {
	w.whtLogic.CreateWht(ctx)
}
