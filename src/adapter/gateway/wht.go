package gateway

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/sky0621/wolf-bff/src/application/domain"

	"golang.org/x/xerrors"

	"github.com/sky0621/wolf-bff/src/adapter/gateway/sqlboilermodel"
	"github.com/sky0621/wolf-bff/src/application"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/sky0621/wolf-bff/src/application/model"
)

func NewWhtRepository(db boil.ContextExecutor) application.WhtRepository {
	return &whtRepository{db: db}
}

type whtRepository struct {
	db boil.ContextExecutor
}

func (r *whtRepository) Create(ctx context.Context, in model.WhtInput) (int64, error) {
	mdl := sqlboilermodel.WHT{
		Recorddate: time.Now(),
		Title:      null.StringFromPtr(in.Title),
	}
	if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
		return -1, xerrors.Errorf("failed to insert wht: %w", err)
	}
	return mdl.ID, nil
}

func (r *whtRepository) Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error) {
	var mod []qm.QueryMod
	if condition != nil {
		if condition.ID != nil {
			mod = append(mod, sqlboilermodel.WHTWhere.ID.EQ(condition.ID.ToPrimitive()))
		}
	}
	records, err := sqlboilermodel.WHTS(mod...).All(ctx, r.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to select wht: %w", err)
	}
	var results []*domain.Wht
	for _, r := range records {
		results = append(results, domain.NewWht(r.ID, r.Recorddate, r.Title.String))
	}
	return results, nil
}
