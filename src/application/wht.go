package application

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/xerrors"

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

func (w Wht) GetWhtByRecordDate(ctx context.Context, recordDate time.Time) (*domain.Wht, error) {
	records, err := w.repository.Read(ctx, &domain.WhtCondition{
		RecordDate: &recordDate,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to repository.Read[recordDate:%#+v]: %w", recordDate, err)
	}
	switch len(records) {
	case 0:
		return nil, nil
	case 1:
		return records[0], nil
	default:
		return nil, xerrors.New(fmt.Sprintf("failed to repository.Read[recordDate:%#+v]: not 1", recordDate))
	}
}

func (w Wht) UpsertWht(ctx context.Context, in domain.Wht) (int64, error) {
	return w.repository.Create(ctx, in)
}

type WhtRepository interface {
	Create(ctx context.Context, in domain.Wht) (int64, error)
	Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
	Upsert(ctx context.Context, in domain.Wht) (*domain.Wht, error)
}
