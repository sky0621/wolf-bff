package application

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application/model"
)

func NewWht(r WhtRepository) *Wht {
	return &Wht{whtRepository: r}
}

type Wht struct {
	whtRepository WhtRepository
}

func (w Wht) CreateWht(ctx context.Context, in model.WhtInput) (string, error) {
	return w.whtRepository.CreateWht(ctx, in)
}

type WhtRepository interface {
	CreateWht(ctx context.Context, in model.WhtInput) (string, error)
}
