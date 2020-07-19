package application

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application/domain"

	"github.com/sky0621/wolf-bff/src/application/model"
)

func NewWht(wht WhtRepository, content ContentRepository) *Wht {
	return &Wht{whtRepository: wht, contentRepository: content}
}

type Wht struct {
	whtRepository     WhtRepository
	contentRepository ContentRepository
}

func (w Wht) CreateWht(ctx context.Context, in model.WhtInput) (int64, error) {
	return w.whtRepository.Create(ctx, in)
}

func (w Wht) ReadWht(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error) {
	return w.whtRepository.Read(ctx, condition)
}

type WhtRepository interface {
	Create(ctx context.Context, in model.WhtInput) (int64, error)
	Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
}
