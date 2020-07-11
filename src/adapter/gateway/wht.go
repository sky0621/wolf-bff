package gateway

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/sky0621/wolf-bff/src/adapter/gateway/sqlboilermodel"
	"github.com/sky0621/wolf-bff/src/application"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/application/model"
)

func NewWhtRepository(db sqlx.Execer) application.WhtRepository {
	return &whtRepository{db: db}
}

type whtRepository struct {
	db sqlx.Execer
}

func (r *whtRepository) CreateWht(ctx context.Context, in model.WhtInput) (int64, error) {
	mdl := sqlboilermodel.WHT{
		Recorddate: time.Now(),
		Title:      null.StringFromPtr(in.Title),
	}
	if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
		return -1, xerrors.Errorf("failed to insert wht: %w", err)
	}
	return mdl.ID, nil
}
