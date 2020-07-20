package application

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application/domain"
)

func NewWht(repository WhtRepository) *Wht {
	return &Wht{repository: repository}
}

type Wht struct {
	repository WhtRepository
}

func (w Wht) CreateWht(ctx context.Context, in domain.Wht) (int64, error) {
	return w.repository.Create(ctx, in)
}

func (w Wht) ReadWht(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error) {
	return w.repository.Read(ctx, condition)
}

type WhtRepository interface {
	Create(ctx context.Context, in domain.Wht) (int64, error)
	Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
}
